package compressor

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
			info, _ := os.Stat(matchFile)
			switch mode := info.Mode(); {
			case mode.IsDir():
				writeDirToCompressedFile(matchFile, zipWriter)
			case mode.IsRegular():
				writeFileToCompressedFile(matchFile, zipWriter)
			}
		}
	}
	return nil
}

func writeDirToCompressedFile(matchFile string, zipWriter *zip.Writer) {
	zipWriter.Create(fmt.Sprintf("%s/", matchFile))
	infos, err := ioutil.ReadDir(matchFile)
	if err != nil {
		return
	}
	for _, info := range infos {
		if !info.IsDir() {
			writeFileToCompressedFile(matchFile+"/"+info.Name(), zipWriter)
		} else {
			writeDirToCompressedFile(matchFile+"/"+info.Name(), zipWriter)
		}
	}
}

func writeFileToCompressedFile(matchFile string, zipWriter *zip.Writer) {
	contents, err := ioutil.ReadFile(matchFile)
	if err != nil {
		return
	}
	writer, err := zipWriter.Create(matchFile)
	if err != nil {
		return
	}
	writer.Write(contents)
}
