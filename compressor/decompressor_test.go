package compressor

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestDecompress(t *testing.T) {

	t.Run("Should decompress given file", func(t *testing.T) {
		hash := "compress"
		rootPath := "../test_resource"
		err := Decompress("../test_resource", hash, rootPath)
		if err != nil {
			t.Error("Expected nil error")
		}

		expectedFiles := []string{"../test_resource/compress",
			"../test_resource/compress/.DS_Store",
			"../test_resource/compress/folders",
			"../test_resource/compress/folders/file",
			"../test_resource/compress/folders/subfolder",
			"../test_resource/compress/folders/subfolder/file"}

		actualFiles := make([]string, 0, 5)

		filesToBeVisited := make([]string, 0, 5)
		filesToBeVisited = append(filesToBeVisited, "../test_resource/compress")

		for len(filesToBeVisited) != 0 {
			for index, fullFileName := range filesToBeVisited {
				actualFiles = append(actualFiles, fullFileName)
				if len(filesToBeVisited) > 1 {
					filesToBeVisited = append(filesToBeVisited[:index], filesToBeVisited[index+1:]...)
				} else {
					filesToBeVisited = []string{}
				}
				info, _ := os.Stat(fullFileName)
				if info.IsDir() {
					infos, _ := ioutil.ReadDir(fullFileName)
					for _, info := range infos {
						filesToBeVisited = append(filesToBeVisited, fullFileName+"/"+info.Name())
					}
				}
			}
		}

		for index, expectedFile := range expectedFiles {
			if expectedFile != actualFiles[index] {
				t.Errorf("Expected %s but was %s", expectedFile, actualFiles[index])
			}
		}
	})
}
