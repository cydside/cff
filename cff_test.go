package cff

import (
	"fmt"
	"testing"
	"time"

	sup "github.com/cydside/cff/src/infrastructure"
)

//______________________________________________________________________________

// TestBoot
// go test -v -run ^TestBoot$
func TestBoot(t *testing.T) {
	sup.Boot()
}

//______________________________________________________________________________

// TestBoot
// go test -v -run ^TestCFIF$
// func TestCFIF(t *testing.T) {
// 	files := CheckFilesInFolder()

// 	for _, v := range files {
// 		dbg.Cyan("%s\n", v)
// 	}
// }

//______________________________________________________________________________

// TestBoot
// go test -v -run ^TestCff$
func TestCff(t *testing.T) {
	f1 := &FolderOptions{
		AbsPath:           "testfolders/a/ah",
		CheckSubfolders:   false,
		IgnoreHiddenFiles: true,
		IntervalCheck:     5,
		CallbackFunction: func(list []string) {
			fmt.Println("f1 found:")
			for _, v := range list {
				fmt.Printf("%s\n", v)
			}
			fmt.Println("")
		},
	}

	f2 := &FolderOptions{
		AbsPath:           "testfolders",
		CheckSubfolders:   true,
		IgnoreHiddenFiles: false,
		IntervalCheck:     2,
		CallbackFunction: func(list []string) {
			fmt.Println("f2 found:")
			for _, v := range list {
				fmt.Printf("%s\n", v)
			}
			fmt.Println("")
		},
	}

	h := New().AddFolder(f1).AddFolder(f2)
	h.Run()
	fmt.Println("_------------------------------")
	time.Sleep(11 * time.Second)

	h.Stop()
}
