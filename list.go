package io

import (
	"io/ioutil"
	"path"
)

// load all files in folderPath
func LoadFilesInFolder(folderPath string) (fileList []string, err error) {
	files, err := ioutil.ReadDir(folderPath)

	if err != nil {
		return
	}

	for _, file := range files {
		if file.IsDir() {
			list, err := LoadFilesInFolder(path.Join(folderPath, file.Name()))
			if err != nil {
				continue
			}
			fileList = append(fileList, list...)
			continue
		}
		fileList = append(fileList, path.Join(folderPath, file.Name()))
	}
	return
}

// load all files in folderPath
//
// It will ignore the hidden files and folders.
func LoadFilesInFolderIgnoreHiddenFiles(folderPath string) (fileList []string, err error) {
	files, err := ioutil.ReadDir(folderPath)

	if err != nil {
		return
	}

	for _, file := range files {
		fileFullPath := path.Join(folderPath, file.Name())
		if yes, err := isHidden(fileFullPath); yes || err != nil {
			continue
		}
		if file.IsDir() {
			list, err := LoadFilesInFolderIgnoreHiddenFiles(fileFullPath)
			if err != nil {
				continue
			}
			fileList = append(fileList, list...)
			continue
		}
		fileList = append(fileList, path.Join(folderPath, file.Name()))
	}
	return
}
