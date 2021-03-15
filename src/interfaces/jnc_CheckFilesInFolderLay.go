package interfaces

//______________________________________________________________________________

// CheckFilesInFolderLay interfaccia
type CheckFilesInFolderLay interface {
	CheckFilesInFolder(absolutePath string) ([]string, error)
	CheckFilesInFolderAndSubfolders(absolutePath string) ([]string, error)
}
