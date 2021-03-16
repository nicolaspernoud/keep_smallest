package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var (
	sourceDir string
	destDir   string
	noDir1    = "doesNotExistDir1"
	noDir2    = "doesNotExistDir2"
)

func TestKeepSmallest(t *testing.T) {
	createTestDirs()
	type args struct {
		sourceDir string
		destDir   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"no_dirs", args{noDir1, noDir2}, true},
		{"same_dirs", args{sourceDir, sourceDir}, true},
		{"correct_dirs", args{sourceDir, destDir}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := KeepSmallest(tt.args.sourceDir, tt.args.destDir); (err != nil) != tt.wantErr {
				t.Errorf("KeepSmallest() error = %v, wantErr %v", err, tt.wantErr)
			} else if err = checkResult(); err != nil && !tt.wantErr {
				t.Errorf("KeepSmallest() error = %v", err)
			}

		})
	}
	deleteTestDirs()
}

func createTestDirs() {
	var err error
	sourceDir, err = ioutil.TempDir("", "sourceDir")
	if err != nil {
		panic(err)
	}
	destDir, err = ioutil.TempDir("", "destDir")
	if err != nil {
		panic(err)
	}
	createTestFile(sourceDir, "file1", 10)
	createTestFile(sourceDir, "file2", 20)
	createTestFile(sourceDir, "file3", 10)
	createTestFile(destDir, "file1", 20)
	createTestFile(destDir, "file2", 10)
	createTestFile(destDir, "file4", 5)
}

func deleteTestDirs() {
	os.RemoveAll(sourceDir)
	os.RemoveAll(destDir)
}

func createTestFile(dir string, name string, size int64) string {
	f, err := os.Create(filepath.Join(dir, name))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := f.Truncate(size); err != nil {
		panic(err)
	}
	return f.Name()
}

func checkResult() error {
	// file1 must still exist in sourceDir
	if f, err := os.Stat(filepath.Join(sourceDir, "file1")); err != nil || f.Size() != 10 {
		return errors.New("file1 must still exist in sourceDir")
	}
	// file1 must exist in destDir, but must be the one that was in sourceDir (size 10)
	if f, err := os.Stat(filepath.Join(destDir, "file1")); err != nil || f.Size() != 10 {
		return errors.New("file1 must exists in destDir, but must be the one that was in sourceDir (size 10)")
	}
	// file2 must still exist in sourceDir
	if f, err := os.Stat(filepath.Join(sourceDir, "file2")); err != nil || f.Size() != 20 {
		return errors.New("file2 must still exist in sourceDir")
	}
	// file2 must exist in destDir and must be the one that was in destDir (size 10)
	if f, err := os.Stat(filepath.Join(destDir, "file2")); err != nil || f.Size() != 10 {
		return errors.New("file2 must exists in destDir and must be the one that was in destDir (size 10)")
	}
	// file3 must exist in sourceDir
	if _, err := os.Stat(filepath.Join(sourceDir, "file3")); err != nil {
		return errors.New("file3 must exist in sourceDir")
	}
	// file3 must not exist in destDir
	if _, err := os.Stat(filepath.Join(destDir, "file3")); err == nil {
		return errors.New("file3 must not exist in destDir")
	}
	// file4 must not exist in sourceDir
	if _, err := os.Stat(filepath.Join(sourceDir, "file4")); err == nil {
		return errors.New("file4 must not exist in sourceDir")
	}
	// file4 must exists in destDir and must be the one that was in destDir (size 5)
	if f, err := os.Stat(filepath.Join(destDir, "file4")); err != nil || f.Size() != 5 {
		return errors.New("file4 must exists in destDir and must be the one that was in destDir (size 5)")
	}
	return nil
}
