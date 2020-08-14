package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	walkDir := "sample"
	toRename := make(map[string][]string)
	filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}
		curDir := filepath.Dir(path)

		if m, err := match(info.Name()); err == nil {
			key := filepath.Join(curDir, fmt.Sprintf("%s.%s", m.base, m.ext))
			toRename[key] = append(toRename[key], info.Name())
		}

		return nil
	})

	// for _, files := range toRename {
	// 	for _, file := range files {
	// 		fmt.Printf("%q\n", file)
	// 	}
	// }

	for key, files := range toRename {
		dir := filepath.Dir(key)
		n := len(files)
		sort.Strings(files)
		for i, filename := range files {
			res, _ := match(filename)
			newFilename := fmt.Sprintf("%s - %d of %d.%s", res.base, (i + 1), n, res.ext)
			oldPath := filepath.Join(dir, filename)
			newPath := filepath.Join(dir, newFilename)
			fmt.Printf("mv %s => %s\n", oldPath, newPath)
			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Println("Error Renaming", oldPath, newPath)
			}
		}
	}
}

type matchResult struct {
	base  string
	ext   string
	index int
}

func match(fileName string) (*matchResult, error) {
	pieces := strings.Split(fileName, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(tmp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return nil, fmt.Errorf("%s did not match our pattern", fileName)
	}

	return &matchResult{strings.Title(name), ext, number}, nil
}
