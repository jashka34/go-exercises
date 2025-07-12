package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func dirTree(out io.Writer, path string, printFiles bool) error {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		fi, err := os.Stat(file.Name())
		if err != nil {
			log.Fatal(err)
		}
		if fi.IsDir() {
			fmt.Println("dir!!!")
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
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
