package io

import (
	"io/ioutil"
	"path"
)

// load all files in folderPath
func loadFilesInFolder(folderPath string) (fileList []string, err error) {
	files, err := ioutil.ReadDir(folderPath)

	if err != nil {
		return
	}

	for _, file := range files {
		if file.IsDir() {
			fileList = append(fileList, loadFilesInFolder(path.Join(folderPath, file.Name()))...)
			continue
		}
		fileList = append(fileList, path.Join(folderPath, file.Name()))
	}
	return
}
