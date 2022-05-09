package main

import (
	"fmt"
	"ioutil"
	"log"
	"strings"
	"sync"
)

var (
	matches []string
	waitgroup = sync.WaitGroup{}
	lock = sync.Mutex{}
)

func search(root string, filename string) {
	fmt.Println("searching in", root)
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal("error on files", err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			lock.Lock()
			matches = append(matches, filepath.Join(root, file.Name())) // c:\tools\Readme.txt
			lock.Unlock()
		}

		if file.IsDir() {
			waitgroup.Add(1)
			go search(filepath.Join(root, file.Name()), filename)
		}
	}
	waitgroup.Done()
}

func main() {
	waitgroup.Add(1)
	go search("C:/tools", "README.md")
	waitgroup.Wait()
	for _, file := range matches {
		fmt.Println("Matched", file)
	}
}