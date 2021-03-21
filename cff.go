package cff

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

//______________________________________________________________________________

// Cff
type Cff struct {
	folders []FolderOptions
	tickers []*time.Ticker
	done    chan bool
	mu      sync.Mutex
}

//______________________________________________________________________________

// https://stackoverflow.com/questions/52381764/run-multiple-functions-concurrently-at-intervals-in-go
// https://play.golang.org/p/L9LiFlc87jo

// FolderOptions
type FolderOptions struct {
	AbsPath           string              `json:"AbsPath"`
	CheckSubfolders   bool                `json:"CheckSubfolders"`
	IgnoreHiddenFiles bool                `json:"IgnoreHiddenFiles"`
	IntervalCheck     int                 `json:"IntervalCheck"`
	CallbackFunction  func(list []string) `json:"-"`
}

//______________________________________________________________________________

func New() *Cff {
	p := &Cff{
		folders: nil,
		tickers: nil,
		done:    make(chan bool),
	}

	return p
}

//______________________________________________________________________________

func (p *Cff) Run() {
	for i := range p.folders {
		t := p.schedule(&p.folders[i], p.done)
		p.tickers = append(p.tickers, t)
	}

}

//______________________________________________________________________________

func (p *Cff) Stop() {
	close(p.done)
	for i := range p.tickers {
		p.tickers[i].Stop()
	}
}

//______________________________________________________________________________

func (p *Cff) AddFolder(obj *FolderOptions) *Cff {
	p.folders = append(p.folders, *obj)

	return p
}

//______________________________________________________________________________

func (p *Cff) schedule(obj *FolderOptions, done <-chan bool) *time.Ticker {
	ticker := time.NewTicker(time.Duration(obj.IntervalCheck) * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				if err := p.checkPoint(obj); err != nil {
					log.Printf("%s\n", err)
					return
				}
				if files, err := p.checkFolder(obj); err != nil {
					log.Printf("%s\n", err)
					return
				} else {
					obj.CallbackFunction(files)
				}
			case <-done:
				return
			}
		}
	}()
	return ticker
}

//______________________________________________________________________________

// checkPoint implementa i controlli obbligatori
func (p *Cff) checkPoint(obj *FolderOptions) error {
	var err error
	var mu sync.Mutex

	mu.Lock()
	defer mu.Unlock()

	if _, err = os.Stat(obj.AbsPath); os.IsNotExist(err) {
		return err
	}

	if obj.IntervalCheck < 5 {
		return errors.New("Interval cannot be < 5 secs")
	}

	return nil
}

//______________________________________________________________________________

// checkFolder implementa
func (p *Cff) checkFolder(obj *FolderOptions) ([]string, error) {
	var err error
	var files []string

	if err = p.checkPoint(obj); err != nil {
		return nil, err
	}

	if obj.CheckSubfolders {
		if files, err = p.getFilesInFolderAndSubfolders(obj.AbsPath); err != nil {
			return nil, err
		}
	} else {
		if files, err = p.getFilesInFolder(obj.AbsPath); err != nil {
			return nil, err
		}
	}

	if obj.IgnoreHiddenFiles {
		return p.removeHiddenFiles(files), nil
	}

	return files, nil
}

//______________________________________________________________________________

func (p *Cff) getFilesInFolder(absolutePath string) ([]string, error) {
	var files []string

	p.mu.Lock()
	defer p.mu.Unlock()

	objs, err := ioutil.ReadDir(absolutePath)
	if err != nil {
		return nil, err
	}

	sort.Slice(objs, func(i, j int) bool { return objs[i].Name() < objs[j].Name() })

	for _, f := range objs {
		if !f.IsDir() {
			// files = append(files, absolutePath+string(os.PathSeparator)+f.Name())
			files = append(files, f.Name())
		}
	}

	return files, nil
}

//______________________________________________________________________________

func (p *Cff) getFilesInFolderAndSubfolders(absolutePath string) ([]string, error) {
	var files []string

	p.mu.Lock()
	defer p.mu.Unlock()

	err := filepath.Walk(absolutePath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	sort.Slice(files, func(i, j int) bool { return files[i] < files[j] })

	return files, nil
}

//______________________________________________________________________________

// removeHiddenFiles rimuove i file nascosti
func (p *Cff) removeHiddenFiles(files []string) []string {
	var mu sync.Mutex
	var r []string

	mu.Lock()
	defer mu.Unlock()

	for _, v := range files {
		if !strings.Contains(v, string(os.PathSeparator)+".") && !strings.HasPrefix(v, ".") {
			r = append(r, v)
		}
	}

	return r
}
