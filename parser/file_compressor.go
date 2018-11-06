package parser

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Compress(fileRegex []string, previousHash string, tempDir string, rootPath string) error {
	newZipFile, err := os.Create(fmt.Sprintf("%s/%s.zip", tempDir, previousHash))
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	for _, file := range fileRegex {
		matches, _ := filepath.Glob(fmt.Sprintf("%s/%s", rootPath, file))
		for _, matchFile := range matches {
			contents, err := ioutil.ReadFile(matchFile)
			if err != nil {
				break
			}
			writer, err := zipWriter.Create(matchFile)
			if err != nil {
				break
			}
			writer.Write(contents)
		}
	}
	return nil
}
