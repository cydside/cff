package domain

//______________________________________________________________________________

// FindInFolder interfaccia
type FindInFolderLay interface {
	Fetch(folderAbsPath string, includeSubfolders bool) ([]string, error)
}
