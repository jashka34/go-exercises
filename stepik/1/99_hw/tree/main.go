package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func dirTree(out io.Writer, path string, printFiles bool) error {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fullPath := filepath.Join(path, file.Name())
		fmt.Println(fullPath)
		fi, err := os.Stat(fullPath)
		if err != nil {
			log.Fatal(err)
		}
		if fi.IsDir() {
			// fmt.Println("dir!!!")
			dirTree(out, fullPath, printFiles)
		}
	}
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	fmt.Println("path: ", path)
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
