package main

import (
	"fmt"
	"ioutil"
	"log"
	"strings"
)

var (
	matches []string
)

func search(root string, filename string) {
	fmt.Println("searching in", root)
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal("error on files", err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			matches = append(matches, filepath.Join(root, file.Name())) // c:\tools\Readme.txt
		}

		if file.IsDir() {
			search(filepath.Join(root, file.Name()), filename)
		}
	}
}

func main() {
	search("C:/tools", "README.md")

	for _, file := range matches {
		fmt.Println("Matched", file)
	}
}