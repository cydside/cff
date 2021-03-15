package infrastructure

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

//______________________________________________________________________________

type CheckFilesInFolderImp struct {
	mu sync.Mutex
}

//______________________________________________________________________________

func (p *CheckFilesInFolderImp) CheckFilesInFolder(absolutePath string) ([]string, error) {
	return p.getFilesInFolder(absolutePath)
}

//______________________________________________________________________________

func (p *CheckFilesInFolderImp) CheckFilesInFolderAndSubfolders(absolutePath string) ([]string, error) {
	return p.getFilesInFolderAndSubfolders(absolutePath)
}

//______________________________________________________________________________

func (p *CheckFilesInFolderImp) getFilesInFolder(absolutePath string) ([]string, error) {
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
			files = append(files, absolutePath+string(os.PathSeparator)+f.Name())
		}
	}

	return files, nil
}

//______________________________________________________________________________

func (p *CheckFilesInFolderImp) getFilesInFolderAndSubfolders(absolutePath string) ([]string, error) {
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

/*
// Versione che NON ricerca nelle sottocartelle
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })

	for _, f := range files {
		if !f.IsDir() {
			fmt.Println(f.Name())
		}
	}
}

// Versione che ricerca nelle sottocartelle
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	var files []string

	root := "."
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	sort.Slice(files, func(i, j int) bool { return files[i] < files[j] })

	for _, file := range files {
		fmt.Println(file)
	}
}
*/
