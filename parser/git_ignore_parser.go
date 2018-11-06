package parser

import (
	"bufio"
	"os"
)

type StringList []string

func Parse(path string) (StringList, error) {
	var contents StringList
	file, e := os.Open(path)
	if e != nil {
		return nil, e
	}
	file.Seek(0,0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}
	return contents, nil
}

func (this StringList) compare(that StringList) bool {
	for index, string := range this {
		if string != that[index] {
			return false
		}
	}
	return true
}