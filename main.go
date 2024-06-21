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

func main() {
	// var findRecursive = flag.Bool("r", false, "Will find in all subdirectories") // Maybe a -d 3 for max depth?
	flag.Parse()
	// args := os.Args
	// fmt.Println(*findRecursive)
	path := flag.Arg(0)

	all, err := DirChildren(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(all)

	err = PrintChildren(all, 0)
	if err != nil {
		log.Fatal(err)
	}
	// for _, file := range all {
	// 	fmt.Println(file, reflect.Slice)

	// 	switch file.(type) {
	// 	case []any:
	// 		fmt.Println(file)
	// 	case fs.DirEntry:
	// 		fmt.Println("a", file)
	// 	}

	// }

	// fmt.Println(flag.Args())
	// files, err = os.ReadDir(path)
}

func PrintChildren(all []any, indent int) error {
	for _, file := range all {
		switch f := file.(type) {
		case fs.DirEntry:
			if f.IsDir() {
				fmt.Println(f.Name(), "/")
			} else {
				fmt.Print(strings.Repeat("  ", indent))
				fmt.Println(f.Name())
			}
		case []any:
			err := PrintChildren(f, indent+1)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func DirChildren(dir_path string) ([]interface{}, error) {
	// var file_list []fs.DirEntry | string
	var file_list []interface{}
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
			file_list = append(file_list, file, dir_c)
			// fmt.Println(dir_path, file_list)
		} else {
			file_list = append(file_list, file)
		}
	}
	return file_list, nil
}
