package scripts

import (
	"fmt"
	"os"
	"path/filepath"
)

//	traverse the specified directory, and for each file or directory in the tree,
//
// it prints whether it's a file or a directory & its path
func ListFilesInDir(dir string) []string {
	var filenames []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fmt.Println("File:", path)
			filenames = append(filenames, path)
		}
		// fmt.Println("Directory:", path)
		return nil
	})
	if err != nil {
		// fmt.Println("Error during filepath.Walk:", err)
		panic(fmt.Errorf("error listing files in %s: %v", dir, err))
	}
	return filenames
}
