package cff

//______________________________________________________________________________

import (
	sup "github.com/cydside/cff/src/infrastructure"
)

//______________________________________________________________________________

func init() {
	sup.Boot()
}

//______________________________________________________________________________

// uiwin
// func CheckFilesInFolder() {
// 	sup.Run()
// }

// uicli
func CheckFilesInFolder() []string {
	return sup.Run()
}
