package copy

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

// File is for copying the file from sourceFolderPath to outputFolderPath
//
// It will return error when the file cannot be copied
func File(sourceFolderPath string, outputFolderPath string, relativeFilePath string) error {
	log.Printf(`Copy file "%s" from "%s" to "%s"`, relativeFilePath, sourceFolderPath, outputFolderPath)
	data, err := ioutil.ReadFile(path.Join(sourceFolderPath, relativeFilePath))
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path.Join(outputFolderPath, relativeFilePath), data, 0644)
}

// FS is for copying the file (path is the FS's path) to outputFolderPath
// It can be copy folder or file
func FS(f embed.FS, path string, outputFolderPath string) {
	fs.WalkDir(f, path, func(filePath string, d fs.DirEntry, err error) error {
		rel, err := filepath.Rel(path, filePath)
		if err != nil {
			return err
		}

		if rel == "." {
			return nil
		}

		if d.IsDir() {
			return os.Mkdir(filepath.Join(outputFolderPath, rel), 0755)
		}

		data, err := f.ReadFile(filePath)
		if err != nil {
			return err
		}
		ioutil.WriteFile(
			filepath.Join(outputFolderPath, rel),
			data,
			0644,
		)

		return nil
	})
}

// Dir is for Copying the source folder to output folder
//
// It will return error when one of sub-file or sub-folder cannot be copied
func Dir(sourceFolder string, outputFolder string) error {
	fileInfos, err := ioutil.ReadDir(sourceFolder)
	if err != nil {
		return fmt.Errorf("Cannot read folder %s", sourceFolder)
	}
	for _, fi := range fileInfos {
		sourcePath := filepath.Join(sourceFolder, fi.Name())
		outPath := filepath.Join(outputFolder, fi.Name())
		if fi.IsDir() {
			err = os.Mkdir(outPath, 0755)
			if err != nil {
				return fmt.Errorf("Cannot create folder %s", outPath)
			}
			err = Dir(sourcePath, outPath)
			if err != nil {
				return err
			}

		} else {
			err = File(sourceFolder, outputFolder, fi.Name())
			if err != nil {
				return err
			}
		}
	}
	return nil
}
