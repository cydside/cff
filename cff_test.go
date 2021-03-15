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
			f1s := []string{"testfolders/a/ah/ah_01.txt"}
			fmt.Println("f1 found:")
			for _, v := range list {
				fmt.Printf("%s\n", v)
			}
			fmt.Println("")
			if !testEq(f1s, list) {
				t.Errorf("Expecting %v, got '%v'\n", f1s, list)
			}
		},
	}

	f2 := &FolderOptions{
		AbsPath:           "testfolders",
		CheckSubfolders:   true,
		IgnoreHiddenFiles: false,
		IntervalCheck:     2,
		CallbackFunction: func(list []string) {
			f1s := []string{
				"testfolders/a/ah/.ah_02.txt",
				"testfolders/a/ah/.ahidden/ahidden_01.txt",
				"testfolders/a/ah/ah_01.txt",
				"testfolders/b/b_01.txt",
				"testfolders/b/boh/boh_01.txt",
				"testfolders/b/boh/boh_02.txt",
				"testfolders/b/boh/boh_03.txt",
			}
			fmt.Println("f2 found:")
			for _, v := range list {
				fmt.Printf("%s\n", v)
			}
			fmt.Println("")
			if !testEq(f1s, list) {
				t.Errorf("Expecting %v, got '%v'\n", f1s, list)
			}
		},
	}

	h := New().AddFolder(f1).AddFolder(f2)
	h.Run()
	time.Sleep(11 * time.Second)
	h.Stop()
}

//______________________________________________________________________________

func testEq(a, b []string) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
