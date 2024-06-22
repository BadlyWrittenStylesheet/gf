package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

var max_depth int
var searchPhrase string
var file_type string
var file_ext string

var active color.Color = *color.New(color.Bold, color.FgHiBlue)
var inactive color.Color = *color.New(color.FgHiBlack, color.Italic)

type FileNode struct {
	fs.DirEntry
	Children []FileNode
}

func main() {
	flag.IntVar(&max_depth, "d", 3, "Set the max depth of file checking.")
	flag.StringVar(&searchPhrase, "n", "", "Search for files with the given phrase in name")
	flag.StringVar(&file_type, "t", "", "Specify if should search for files or directories")
	flag.StringVar(&file_ext, "e", "", "Search for files with a given extension")

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
	// fmt.Println(all)
}

func PrintChildren(all []FileNode, indent int) error {
	for i, file := range all {

		var ind string
		if indent != 0 {
			if i == len(all)-1 {
				ind = strings.Repeat("  ", max(indent)) + "└-- "
			} else {
				ind = strings.Repeat("  ", max(indent)) + "|-- "
			}
		}

		fmt.Print(ind)
		if len(file.Children) == 0 || strings.Contains(file.Name(), searchPhrase) {
			active.Print(file.Name())
		} else {
			inactive.Print(file.Name())
		}
		if file.IsDir() {
			fmt.Println("/")
			err := PrintChildren(file.Children, indent+1)
			if err != nil {
				return err
			}
		} else {
			fmt.Print("\n")
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
	return filterByName(file) && filterByExtension(file) && filterByType(file)
}

func shouldIncludeDir(file fs.DirEntry, children []FileNode) bool {
	return (filterByName(file) && filterByType(file)) || len(children) != 0
}

func filterByName(file fs.DirEntry) bool {
	if file_ext != "" && file.IsDir() {
		return false
	}
	return searchPhrase == "" || strings.Contains(file.Name(), searchPhrase)
}

func filterByType(file fs.DirEntry) bool {
	// log.Println(file, file_type, file.IsDir())
	switch file_type {
	case "d":
		return file.IsDir()
	case "f":
		return !file.IsDir()
	default:
		return true
	}
}

func filterByExtension(file fs.DirEntry) bool {
	return file_ext == "" || strings.HasSuffix(
		file.Name(), "."+
			strings.TrimPrefix(
				file_ext, "."))
}
