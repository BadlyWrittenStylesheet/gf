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
var searchPhrase string

// var file_type string

var file_ext string

type FileNode struct {
	fs.DirEntry
	Children []FileNode
}

func main() {
	// var findRecursive = flag.Bool("r", false, "Will find in all subdirectories") // Maybe a -d 3 for max depth?
	flag.IntVar(&max_depth, "d", 3, "Set the max depth of file checking.")
	flag.StringVar(&searchPhrase, "n", "", "Search for files with the given phrase in name")
	flag.StringVar(&file_ext, "e", "", "Search for files with a given extension")
	// flag.StringVar(&file_type, "t", "", "Specify if should search for files or directories")

	flag.Parse()
	// fmt.Println(max_depth, searchPhrase)

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
	// fmt.Println(all)
}

func PrintChildren(all []FileNode, indent int) error {
	for _, file := range all {

		fmt.Print(strings.Repeat("  ", indent))
		if file.IsDir() {
			fmt.Printf("%s/\n", file.Name())
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
		if current_depth <= max_depth {
			if file.IsDir() {

				dirChildren, err := GetDirectoryChildren(dir_path+"/"+file.Name(), current_depth+1)
				if err != nil {
					return nil, err
				}

				// log.Println("File: ", file, "children: ", dirChildren, shouldIncludeDir(file, dirChildren))
				if shouldIncludeDir(file, dirChildren) {
					file_list = append(file_list, FileNode{file, dirChildren})
				}
			} else {

				if shouldIncludeFile(file) {
					file_list = append(file_list, FileNode{file, nil})
				}
			}
		}
	}
	return file_list, nil
}

func shouldIncludeFile(file fs.DirEntry) bool {
	return filterByName(file) && filterByExtension(file)
}

func shouldIncludeDir(file fs.DirEntry, children []FileNode) bool {
	return filterByName(file) || len(children) != 0
}

func filterByName(file fs.DirEntry) bool {
	if file_ext != "" && file.IsDir() {
		return false
	}
	return searchPhrase == "" || strings.Contains(file.Name(), searchPhrase)
}

func filterByExtension(file fs.DirEntry) bool {
	return file_ext == "" || strings.HasSuffix(file.Name(), "."+file_ext)
}
