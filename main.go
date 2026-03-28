package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	fmt.Println("Treemaker v1.0 — starting…")
	time.Sleep(3 * time.Second) // creates delay before the next line
	fmt.Println()
	fmt.Println()
	fmt.Println("Please kindly paste/write you're project tree")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	var pathStack []string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == ""{
			fmt.Printf("\nProcessing tree...\n")
			break // stops processing on an empty line
		}
		
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
		// fmt.Println("Full path:", FullPath)

		if isDir {
			err := os.MkdirAll(FullPath, 0755)
			if err != nil {
				fmt.Println("Error Creating directory: ", err)
				return
			}
			fmt.Println("\nCreated directory: ", FullPath)
		}else{
			dir := filepath.Dir(FullPath)
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				fmt.Println("Error creating parent directories: ", err)
				continue
			}

			f, err := os.Create(FullPath)
			if err != nil {
				fmt.Println("Error creating file: ", err)
				continue
			}
			f.Close()
			fmt.Println("\nCreated file: ", FullPath)
		}
	}
	fmt.Println("\n Project tree successfully created!")
	
	if err := scanner.Err(); err != nil{
		fmt.Println("Error reading input:", err)
	}
}