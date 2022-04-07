package copy

import (
	"embed"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

// Copy the file from sourceFolderPath to outputFolderPath
func File(sourceFolderPath string, outputFolderPath string, relativeFilePath string) {
	log.Printf(`Copy file "%s" from "%s" to "%s"`, relativeFilePath, sourceFolderPath, outputFolderPath)
	data, err := ioutil.ReadFile(path.Join(sourceFolderPath, relativeFilePath))
	if err != nil {
		return
	}

	ioutil.WriteFile(path.Join(outputFolderPath, relativeFilePath), data, 0644)
}

// Copy the file (path is the FS's path) to outputFolderPath
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

//Copy the source folder to output folder
func Dir(sourceFolder string, outputFolder string) {
	fileInfos, err := ioutil.ReadDir(sourceFolder)
	if err != nil {
		log.Printf("Cannot read folder %s", sourceFolder)
	}
	for _, fi := range fileInfos {
		sourcePath := filepath.Join(sourceFolder, fi.Name())
		outPath := filepath.Join(outputFolder, fi.Name())
		if fi.IsDir() {
			err = os.Mkdir(outPath, 0755)
			if err != nil {
				log.Printf("Cannot create folder %s", outPath)
			} else {
				copyDir(sourcePath, outPath)
			}

		} else {
			copyFile(sourcePath, outputFolder, fi.Name())
		}
	}
}
