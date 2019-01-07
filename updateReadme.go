package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	f, err := os.OpenFile("README.md", os.O_APPEND|os.O_WRONLY, 0600)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	names, err := getFilesNames("golang/leetcode")
	if err != nil {
		panic(err)
	}

	for _, n := range names {
		if _, err = f.WriteString(n + "\n"); err != nil {
			panic(err)
		}
	}


}

func getFilesNames(dir string) ([]string, error) {
	var filesNames []string

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		name := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
		filesNames = append(filesNames, name)
	}

	return filesNames, nil
}