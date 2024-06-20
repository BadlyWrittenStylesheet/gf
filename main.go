package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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

	fmt.Println(flag.Args())
	// files, err = os.ReadDir(path)
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
			file_list = append(file_list, dir_c)
			fmt.Println(dir_path, file_list)
		} else {
			file_list = append(file_list, file)
		}
	}
	return file_list, nil
}
