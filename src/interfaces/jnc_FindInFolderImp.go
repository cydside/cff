package interfaces

//______________________________________________________________________________

// FindInFolderImp implementa FindInFolderLay
type FindInFolderImp struct {
	CheckFilesInFolderLay CheckFilesInFolderLay
}

//______________________________________________________________________________

// NewFindInFolderImp new
func NewFindInFolderImp(ad CheckFilesInFolderLay) *FindInFolderImp {
	k := new(FindInFolderImp)
	k.CheckFilesInFolderLay = ad

	return k
}

//______________________________________________________________________________

// Fetch implementa
func (p *FindInFolderImp) Fetch(folderAbsPath string, includeSubfolders bool) ([]string, error) {
	if includeSubfolders {
		return p.CheckFilesInFolderLay.CheckFilesInFolderAndSubfolders(folderAbsPath)
	} else {
		return p.CheckFilesInFolderLay.CheckFilesInFolder(folderAbsPath)
	}
}
