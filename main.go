package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

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

func dirTree(output io.Writer, thisPath string, showFiles bool) error {

	os.Chdir(thisPath)
	prefixStart := "├───"
	printDir(prefixStart, showFiles, output)

	return nil

}

func printDir(prefixStart string, showFiles bool, output io.Writer) {
	dirList, _ := os.ReadDir(".")
	dirListFoldersOnly := make([]os.DirEntry, 0)
	if !showFiles {
		for j := range dirList {
			if dirList[j].IsDir() {
				dirListFoldersOnly = append(dirListFoldersOnly, dirList[j])
			}
		}
		dirList = dirListFoldersOnly
	}

	for i := range dirList {
		prefix := prefixStart
		if i == len(dirList)-1 {
			prefix = prefix[:len(prefix)-12] + "└───"
		} else {
			prefix = prefix[:len(prefix)-12] + "├───"
		}
		if !dirList[i].IsDir() && showFiles {
			fileInfo, _ := dirList[i].Info()

			fileSizeStr := "(empty)"

			if fileInfo.Size() != 0 {
				fileSizeStr = "(" + strconv.FormatInt((fileInfo.Size()), 10) + "b)"
			}

			fmt.Fprintln(output, prefix+dirList[i].Name()+" "+fileSizeStr)
		} else if dirList[i].IsDir() {
			fmt.Fprintln(output, prefix+dirList[i].Name())
		}

		if dirList[i].IsDir() {
			os.Chdir(dirList[i].Name())
			if i == len(dirList)-1 {
				prefix = prefix[:len(prefix)-12] + "\t└───"
			} else {
				prefix = prefix[:len(prefix)-12] + "│\t├───"
			}
			printDir(prefix, showFiles, output)
			os.Chdir("..")

		}
	}
}
