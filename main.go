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
var file_name string

// var file_type string
// var file_ext string

type FileNode struct {
	fs.DirEntry
	Children []FileNode
}

func main() {
	// var findRecursive = flag.Bool("r", false, "Will find in all subdirectories") // Maybe a -d 3 for max depth?
	flag.IntVar(&max_depth, "d", 3, "Set the max depth of file checking.")
	flag.StringVar(&file_name, "n", "", "Search for files with the given phrase in name")
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
				// this part skips all directtories that have no children if a file_name is specified and the directory doesn't contain that phrase
				if file_name != "" && len(dir_c) == 0 && !strings.Contains(file.Name(), file_name) {
					continue
				}
				file_list = append(file_list, FileNode{file, dir_c})
			} else {
				file_list = append(file_list, FileNode{file, nil})
			}
		} else {

			// this part is responsible for only appending files that contain the file_name content in them
			if file_name != "" {
				if strings.Contains(file.Name(), file_name) {
					file_list = append(file_list, FileNode{file, nil})
				}
			} else {
				file_list = append(file_list, FileNode{file, nil})
			}
		}
	}
	return file_list, nil
}
