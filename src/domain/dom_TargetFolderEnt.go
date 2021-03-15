package domain

//______________________________________________________________________________

// TargetFolder configuration
type TargetFolder struct {
	AbsPath           string `json:"AbsPath"`
	CheckSubfolders   bool   `json:"CheckSubfolders"`
	IgnoreHiddenFiles bool   `json:"IgnoreHiddenFiles"`
}
