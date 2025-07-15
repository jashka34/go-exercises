package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func dirTree2(out io.Writer, path string, printFiles bool, indent int) error {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fullPath := filepath.Join(path, file.Name())
		fi, err := os.Stat(fullPath)
		fmt.Println(indent, "->", fullPath, "(", fi.Size(), ")")
		// out.Write([]byte(fullPath + "\n"))
		if err != nil {
			log.Fatal(err)
		}
		if fi.IsDir() {
			// fmt.Println("dir!!!")
			dirTree2(out, fullPath, printFiles, indent+1)
		}
	}
	return nil
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	return dirTree2(out, path, printFiles, 1)
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
