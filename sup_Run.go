package infrastructure

import (
	jnc "github.com/cydside/cff/src/interfaces"
	ui "github.com/cydside/cff/src/interfaces/uicli"

	// ui "github.com/cydside/cff/src/interfaces/uiwin"
	usc "github.com/cydside/cff/src/usecases"
)

//______________________________________________________________________________

// Run lancia l'applicazione
func Run() []string {
	x := new(CheckFilesInFolderImp)

	u := &usc.CheckFolderImp{}
	u.FindInFolderLay = jnc.NewFindInFolderImp(x)

	s := new(ui.MainUICtr)
	s.CheckFolderLay = u

	// s.Land(AppConfig.Folders)

	return removeDupsInStringSlice(s.Land(AppConfig.Folders))
}
