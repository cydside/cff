package cff

import (
	"fmt"
	"testing"
	"time"
)

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
	time.Sleep(11 * time.Second)
	h.Stop()
}
