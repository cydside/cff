package cff

import (
	"testing"

	sup "github.com/cydside/cff/src/infrastructure"
	dbg "github.com/fatih/color"
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
func TestCFIF(t *testing.T) {
	files := CheckFilesInFolder()

	for _, v := range files {
		dbg.Cyan("%s\n", v)
	}
}
