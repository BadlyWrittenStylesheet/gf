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

var max_depth int

type FileNode struct {
	fs.DirEntry
	Children []FileNode
}

func main() {
	// var findRecursive = flag.Bool("r", false, "Will find in all subdirectories") // Maybe a -d 3 for max depth?
	flag.IntVar(&max_depth, "d", 3, "Set the max depth of file checking.")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("Where do you want to search bruh?")
	}
	path := args[0]

	all, err := GetDirectoryChildren(path, 1)
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

func GetDirectoryChildren(dir_path string, current_depth int) ([]FileNode, error) {
	var file_list []FileNode
	dir, err := os.ReadDir(dir_path)
	if err != nil {
		return nil, err
	}
	for _, file := range dir {
		if file.IsDir() {
			if current_depth < max_depth {
				dir_c, err := GetDirectoryChildren(dir_path+"/"+file.Name(), current_depth+1)
				if err != nil {
					return nil, err
				}
				file_list = append(file_list, FileNode{file, dir_c})
			} else {
				file_list = append(file_list, FileNode{file, nil})
			}
		} else {
			file_list = append(file_list, FileNode{file, nil})
		}
	}
	return file_list, nil
}
