package compressor

import (
	"archive/zip"
	"fmt"
	"testing"
)

func TestCompress(t *testing.T) {
	t.Run("Should compress all the files given in the arg", func(t *testing.T) {
		fileRegex := []string{".idea", "normalfile"}
		previousHash := "hashfortest1"
		tempDir := "../test_resource"
		rootPath := "../test_resource"

		err := Compress(fileRegex, previousHash, tempDir, rootPath)

		if err != nil {
			t.Error("Expected error to be nil")
			return
		}

		r, err := zip.OpenReader(fmt.Sprintf("%s/%s.zip", tempDir, previousHash))
		defer r.Close()

		if err != nil {
			t.Error("Expected error to be nil on opening the output file")
		}

		if len(r.File) != 2 {
			t.Errorf("Expected 2 files but found %d", len(r.File))
		}

		for index, file := range r.File {
			expectedFileName := fmt.Sprintf("%s/%s", rootPath, fileRegex[index])
			if file.Name != expectedFileName {
				t.Errorf("Expected file %s but was %s", expectedFileName, file.Name)
			}
		}

	})

	t.Run("Should compress all the directories present in the path", func(t *testing.T) {
		fileRegex := []string{"*dir"}
		previousHash := "hashfortest2"
		tempDir := "../test_resource"
		rootPath := "../test_resource"

		err := Compress(fileRegex, previousHash, tempDir, rootPath)

		if err != nil {
			t.Error("Expected error to be nil")
			return
		}

		r, err := zip.OpenReader(fmt.Sprintf("%s/%s.zip", tempDir, previousHash))
		defer r.Close()

		if err != nil {
			t.Error("Expected error to be nil on opening the output file")
		}

		if len(r.File) != 6 {
			t.Errorf("Expected 2 files but found %d", len(r.File))
		}

		expectedFileNames := []string{
			"../test_resource/.dir/", "../test_resource/.dir/fileinside_dir",
			"../test_resource/another dir/", "../test_resource/another dir/file",
			"../test_resource/yet_another_dir/", "../test_resource/yet_another_dir/file"}

		for index, file := range r.File {
			if file.Name != expectedFileNames[index] {
				t.Errorf("Expected file %s but was %s", expectedFileNames[index], file.Name)
			}
		}
	})
}
