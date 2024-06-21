package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	// "os"
	// "flag"
)

type FileNode struct {
	fs.DirEntry
	Children []FileNode
}

func main() {
	// var findRecursive = flag.Bool("r", false, "Will find in all subdirectories") // Maybe a -d 3 for max depth?
	flag.Parse()

	path := flag.Arg(0)

	all, err := DirChildren(path)
	if err != nil {
		log.Fatal(err)
	}

	err = PrintChildren(all, 0)
	if err != nil {
		log.Fatal(err)
	}
}

func PrintChildren(all []FileNode, indent int) error {
	for _, file := range all {

		fmt.Print(strings.Repeat("  ", indent))
		if file.IsDir() {
			fmt.Print(file.Name())
			fmt.Println("/")
			err := PrintChildren(file.Children, indent+1)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(file.Name())
		}
	}
	return nil
}

func DirChildren(dir_path string) ([]FileNode, error) {
	var file_list []FileNode
	dir, err := os.ReadDir(dir_path)
	if err != nil {
		return nil, err
	}
	for _, file := range dir {
		if file.IsDir() {
			dir_c, err := DirChildren(dir_path + "/" + file.Name())
			if err != nil {
				return nil, err
			}
			file_list = append(file_list, FileNode{file, dir_c})
		} else {
			file_list = append(file_list, FileNode{file, nil})
		}
	}
	return file_list, nil
}
