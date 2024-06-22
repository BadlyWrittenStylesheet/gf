package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/fatih/color"
)

var (
	max_depth       int
	searchPhrase    string
	searchType      string
	searchExtension string
	sortFiles       string
	regexpPhrase    string
	regexFilter     *regexp.Regexp
)

var active color.Color = *color.New(color.Bold, color.FgBlue)
var inactive color.Color = *color.New(color.FgHiBlack, color.Italic)

type FileNode struct {
	fs.DirEntry
	Children []FileNode
}

func main() {
	flag.IntVar(&max_depth, "d", 3, "Set the max depth of file checking.")
	flag.StringVar(&searchPhrase, "n", "", "Search for files with the given phrase in name")
	flag.StringVar(&searchType, "t", "", "Specify if should search for files or directories")
	flag.StringVar(&searchExtension, "e", "", "Search for files with a given extension")
	flag.StringVar(&sortFiles, "s", "", "Sort files by name, size or modification time")
	flag.StringVar(&regexpPhrase, "x", "", "Search name with regular expressions")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("Where do you want to search bruh?")
	}
	if regexpPhrase != "" {
		var err error
		regexFilter, err = regexp.Compile(regexpPhrase)
		if err != nil {
			log.Fatal(err)
		}
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
	for i, file := range all {

		var ind string = strings.Repeat("  ", max(indent))
		if indent != 0 {
			if i == len(all)-1 {
				ind += " └-"
			} else {
				ind += " ├-"
			}
			if len(file.Children) != 0 {
				ind += "┬"
			} else {
				ind += " "
			}
			ind += " "
		}

		fmt.Print(ind)

		if !file.IsDir() && ((searchPhrase != "" &&
			strings.Contains(file.Name(), searchPhrase)) ||
			(filterByExtension(file))) {
			active.Print(file.Name())
		} else if filterByName(file) {
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
				switch sortFiles {
				case "name":
					sortFilesByName(dirChildren)
				case "size":
					sortFilesBySize(dirChildren)
				case "date":
					sortFilesByDate(dirChildren)
				}
				if err != nil {
					return nil, err
				}

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
	if searchExtension != "" && file.IsDir() {
		return false
	}
	return (searchPhrase == "" || strings.Contains(file.Name(), searchPhrase)) &&
		filterNameByRegex(file.Name())
}

func filterByType(file fs.DirEntry) bool {
	// log.Println(file, searchType, file.IsDir())
	switch searchType {
	case "d":
		return file.IsDir()
	case "f":
		return !file.IsDir()
	default:
		return true
	}
}

func filterByExtension(file fs.DirEntry) bool {
	return (searchExtension == "" ||
		strings.HasSuffix(
			file.Name(), "."+strings.TrimPrefix(searchExtension, ".")))
	// regexpPhrase == "" ||
	// filterNameByRegex(file.Name())
}

func filterNameByRegex(name string) bool {
	if regexFilter != nil {
		return (*regexFilter).MatchString(name)
	}
	return true
}

func sortFilesByName(files []FileNode) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
}

func sortFilesBySize(files []FileNode) {
	sort.Slice(files, func(i, j int) bool {
		infoI, _ := files[i].Info()
		infoJ, _ := files[j].Info()
		return infoI.Size() < infoJ.Size()
	})
}

func sortFilesByDate(files []FileNode) {
	sort.Slice(files, func(i, j int) bool {
		infoI, _ := files[i].Info()
		infoJ, _ := files[j].Info()
		return infoI.ModTime().Before(infoJ.ModTime())
	})

}
