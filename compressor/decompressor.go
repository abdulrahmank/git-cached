package compressor

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func Decompress(tmpFilePath string, hash string, rootPath string) error {
	reader, e := zip.OpenReader(fmt.Sprintf("%s/%s.zip", tmpFilePath, hash))
	if e != nil {
		log.Fatalf("%s", e.Error())
	}
	defer reader.Close()
	for _, file := range reader.File {
		f, _ := file.Open()
		if !file.Mode().IsDir() {
			file, err := os.Create(rootPath + "/" + file.Name)
			if err != nil {
				log.Fatalf("Error while creating file: %s", err.Error())
			}
			io.Copy(file, f)
		} else {
			err := os.RemoveAll(rootPath + "/" + file.Name)
			if err != nil {
				log.Fatalf("Unable to remove exisitng ignore files: %s", err.Error())
			}
			os.MkdirAll(rootPath+"/"+file.Name, os.ModePerm)
		}
	}
	return nil
}
