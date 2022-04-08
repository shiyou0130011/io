package copy

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestCopyFS(t *testing.T) {
	m := fstest.MapFS{
		"example/foo.txt": {
			Data: []byte("Nunc aliquam orci eget tortor tristique interdum."),
			Mode: 0644,
		},
		"example/bar.txt": {
			Data: []byte("Proin venenatis urna et velit feugiat euismod."),
			Mode: 0644,
		},
		"foobar": {
			Data: []byte("Curabitur posuere lacus non quam suscipit vestibulum."),
			Mode: 0644,
		},
	}

	outputFolder, err := ioutil.TempDir(os.TempDir(), "test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(outputFolder)
	FS(m, "example", outputFolder)

	fooData, err := ioutil.ReadFile(filepath.Join(outputFolder, "foo.txt"))
	if err != nil {
		t.Fatal(err)
	}

	if string(fooData) != string(m["example/foo.txt"].Data) {
		t.Fatal("foo.txt is not the same")
	}
	t.Logf("Compare %s successfully", "foo.txt")

	barData, err := ioutil.ReadFile(filepath.Join(outputFolder, "bar.txt"))
	if err != nil {
		t.Fatal(err)
	}

	if string(barData) != string(m["example/bar.txt"].Data) {
		t.Fatal("bar.txt is not the same")
	}
	t.Logf("Compare %s successfully", "bar.txt")
}
