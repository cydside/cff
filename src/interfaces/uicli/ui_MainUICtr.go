package uicli

import (
	dbg "github.com/fatih/color"

	dom "github.com/cydside/cff/src/domain"
)

//______________________________________________________________________________

// MainUICtr interagisce
type MainUICtr struct {
	CheckFolderLay CheckFolderLay
}

//______________________________________________________________________________

// Land handler
func (p *MainUICtr) Land(folders []dom.TargetFolder) []string {
	var allFiles []string

	for _, v := range folders {
		var err error
		var files []string

		if files, err = p.CheckFolderLay.CheckFolder(&v); err != nil {
			dbg.Red("Error checking folder \"%s\":\n%s\n", v.AbsPath, err)
			continue
		}

		for _, f := range files {
			dbg.Blue("Found: %s\n", f)
		}

		allFiles = append(allFiles, files...)
	}

	return allFiles
}
