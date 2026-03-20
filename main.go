package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var pathStack []string

	for scanner.Scan() {
		line := scanner.Text()
		
		trimmed := strings.TrimLeft(line, " ")
		indent := len(line) - len(trimmed)
		level := indent / 2 
		
		isDir := strings.HasSuffix(trimmed, "/")
		name := strings.TrimSuffix(trimmed, "/")



		
		if level < len(pathStack) {
			pathStack = pathStack[:level]
		}
		
		pathStack = append(pathStack, name)
		
		
		FullPath := filepath.Join(pathStack...)
		fmt.Println("Full path:", FullPath)

		if isDir == true {
			os.MkdirAll(FullPath, 0755)
		}else{
			f, err := os.Create(FullPath)
			if err != nil {
				defer f.Close()
			}
		}
		
		err := os.MkdirAll(FullPath, 0755)
		if err != nil {
			fmt.Println("Error creating directory: ", err)
		}
	}
	
	if err := scanner.Err(); err != nil{
		fmt.Println("Error reading input:", err)
	}
}