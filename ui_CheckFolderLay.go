package uiwin

//______________________________________________________________________________

import (
	dom "github.com/cydside/cff/src/domain"
)

//______________________________________________________________________________

// CheckFolderLay interfaccia
type CheckFolderLay interface {
	CheckFolder(obj *dom.TargetFolder) ([]string, error)
}
