package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
)

func Test_generateReport(t *testing.T) {
	fs := afero.NewMemMapFs()
	wd, _ := os.Getwd()
	filename := filepath.Join(wd, "test.md")

	af := afero.Afero{Fs: fs}
	ext, _ := af.Exists(filename)
	if ext {
		t.Fatal(filename, " already exists")
	}
	_ = generateReport(filename, fs)
	ext, _ = af.Exists(filename)
	if !ext {
		t.Fatal(filename, " not found")
	}
}
