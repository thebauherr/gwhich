package main

import (
	"fmt"
	"os"
  "path/filepath"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("ERROR: No arguments, provide a binary name!")
		return
	} else if len(args) > 2 {
		fmt.Println("WARNING: This program doesn't support multiple arguments!")
		return
	}

	path := os.Getenv("PATH")
	pathList := filepath.SplitList(path)

	for _, directory := range pathList {
		fullPath := filepath.Join(directory, args[1])
		fileDesc, err := os.Stat(fullPath)

		if err == nil {
			mode := fileDesc.Mode()
			if mode.IsRegular() && mode&0111 != 0 {
				fmt.Println(fullPath)
				return
			}
		}
	}
}
