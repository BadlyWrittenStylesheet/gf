package main

import (
	"flag"
	"fmt"
	"io/fs"
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

	dir, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range dir {
		// fmt.Println(file.Name(), file.IsDir(), file.Type())
		if file.IsDir() {
			fmt.Println("Directory: ", file.Name())
			dir, err = os.ReadDir(path + "/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(dir)
		} else {
			fmt.Println("File: ", file.Name())
		}
	}

	fmt.Println(flag.Args())
	// files, err = os.ReadDir(path)
}

func DirChildren(dir_path string) ([]fs.DirEntry, error) {
	dir, err := os.ReadDir(dir_path)

	if err != nil {
		return nil, err
	}
	return dir, nil
}
