package usecases

import (
	"os"
	"strings"
	"sync"

	dom "github.com/cydside/cff/src/domain"
)

//______________________________________________________________________________

type CheckFolderImp struct {
	FindInFolderLay dom.FindInFolderLay
}

//______________________________________________________________________________

func (p *CheckFolderImp) CheckFolder(obj *dom.TargetFolder) ([]string, error) {
	var err error
	var files []string

	if err = p.checkPoint(obj); err != nil {
		return nil, err
	}

	if files, err = p.FindInFolderLay.Fetch(obj.AbsPath, obj.CheckSubfolders); err != nil {
		return nil, err
	}

	if obj.IgnoreHiddenFiles {
		return p.removeHiddenFiles(files), nil
	}

	return files, nil
}

//______________________________________________________________________________

// checkPoint implementa i controlli obbligatori
func (p *CheckFolderImp) checkPoint(obj *dom.TargetFolder) error {
	var err error
	var mu sync.Mutex

	mu.Lock()
	defer mu.Unlock()

	if _, err = os.Stat(obj.AbsPath); os.IsNotExist(err) {
		return err
	}

	return nil
}

//______________________________________________________________________________

// removeHiddenFiles rimuove i file nascosti
func (p *CheckFolderImp) removeHiddenFiles(files []string) []string {
	var mu sync.Mutex
	var r []string

	mu.Lock()
	defer mu.Unlock()

	for _, v := range files {
		if !strings.Contains(v, string(os.PathSeparator)+".") {
			r = append(r, v)
		}
	}

	return r
}
