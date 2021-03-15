package uiwin

import (
	g "github.com/AllenDang/giu"
	dom "github.com/cydside/cff/src/domain"
	dbg "github.com/fatih/color"
)

//______________________________________________________________________________

// MainUICtr interagisce
type MainUICtr struct {
	CheckFolderLay CheckFolderLay
	files          []string
	folders        []string
}

//______________________________________________________________________________

// Land handler
func (p *MainUICtr) Land(folders []dom.TargetFolder) {
	for _, v := range folders {
		var err error
		var files []string
		p.folders = append(p.folders, v.AbsPath)

		if files, err = p.CheckFolderLay.CheckFolder(&v); err != nil {
			dbg.Red("Error checking folder \"%s\":\n%s\n", v.AbsPath, err)
			continue
		}

		for _, f := range files {
			dbg.Blue("Found: %s\n", f)
		}

		p.files = append(p.files, files...)
	}

	wnd := g.NewMasterWindow("Huge list demo", 600, 310, 0, nil)
	wnd.Run(p.loop)
}

//______________________________________________________________________________

func (p *MainUICtr) loop() {
	g.SingleWindow("Huge list demo").Layout(
		g.Child("FilterContent1").Border(false).Size(0, 150).Layout(
			g.Label("Note: FastTable only works if all rows have same height").Wrapped(true),
			g.ListBox("ListBox1", p.folders),
		),
		g.Child("FilterContent").Border(true).Size(0, 150).Layout(
			g.Label("Note: FastTable only works if all rows have same height").Wrapped(true),
			g.ListBox("ListBox2", p.files),
		),
	)
}
