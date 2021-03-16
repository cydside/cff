# cff
A Golang Check Files in Folder package. It provides options to search or not in subfolders, include or exclude hidden files, interval for periodic check and a callback function support.

## Usage examples:

```
	f1 := &FolderOptions{
		AbsPath:           "testfolders/a/ah",
		CheckSubfolders:   false,
		IgnoreHiddenFiles: true,
		IntervalCheck:     5,
		CallbackFunction: func(list []string) {
			fmt.Println("f1 has found:")
			for _, v := range list {
				fmt.Printf("%s\n", v)
			}
			fmt.Println("")
		},
	}

	f2 := &FolderOptions{
		AbsPath:           "testfolders",
		CheckSubfolders:   true,
		IgnoreHiddenFiles: false,
		IntervalCheck:     8,
		CallbackFunction: func(list []string) {
			fmt.Println("f2 has found:")
			for _, v := range list {
				fmt.Printf("%s\n", v)
			}
			fmt.Println("")
		},
	}

	h := New().AddFolder(f1).AddFolder(f2)
	h.Run()
	time.Sleep(11 * time.Second)
	h.Stop()
```

Feel free to suggest improvements.
