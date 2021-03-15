package cff

import (
	"testing"

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
func TestCFIF(t *testing.T) {
	CheckFilesInFolder()
}
