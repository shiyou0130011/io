package copy_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/shiyou0130011/io/copy"
)

func Example_copyDir() {
	sourceFolder, err := ioutil.TempDir(os.TempDir(), "source-*")
	if err != nil {
		log.Println("Cannot create source folder")
	}
	defer os.RemoveAll(sourceFolder)
	targetFolder, err := ioutil.TempDir(os.TempDir(), "target-*")
	if err != nil {
		log.Fatal("Cannot create target folder")
	}
	defer os.RemoveAll(targetFolder)

	ioutil.WriteFile(filepath.Join(sourceFolder, "foo.txt"), []byte("Foo"), 0644)
	ioutil.WriteFile(filepath.Join(sourceFolder, "bar.txt"), []byte("Foo"), 0644)

	err = copy.Dir(sourceFolder, targetFolder)
	if err != nil {
		log.Fatal(err)
	}
	files, err := os.ReadDir(targetFolder)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Target folder has %d files", len(files))

	// Output:
	// Target folder has 2 files
}
