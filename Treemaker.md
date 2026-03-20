# GUIDE ON HOW TO BUILD A TREEMAKER

---
### ***⚠️🛑 Complete The Exercise Before Checking The Correct Code Below The Exercise.***
---

## 🟢 Phase 1: Read input line by line

#### Goal: 
Learn to read a multi-line tree from stdin.

### Exercise 1:

#### Write a Go program that:

Reads lines from standard input (os.Stdin)

Prints each line prefixed with "Line: "

#### Hints:

Use ```bufio.NewScanner(os.Stdin) ```

Use a for ```scanner.Scan()``` loop

Use ```scanner.Text()``` to get each line

#### Example input (what you’ll paste after running go run main.go):
```
project/
  cmd/
    main.go
```
#### Expected output:
```bash
Line: project/
Line:   cmd/
Line:     main.go
```
##### 💡 Task: Write this code, run it, and paste your working snippet here.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()     
		fmt.Println("Line:", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
```


### ⚡ Mini Exercise (practice)

Change the program so that each line is printed in uppercase instead of "Line: ...".

#### Hint: 
use ```strings.ToUpper(line)``` (Study the strings package).

***Write your version and test it.***

## 🟢 Phase 2: Calculate indentation (tree depth)

Next, we learn how to know which line is a folder/file and its depth.

#### Goal:

* Count leading spaces for each line

* Divide by 2 → depth level

* Print both line text and its level

### Exercise 2:

#### Modify your program so it prints:
```bash
Line: project/   Level: 0
Line:   cmd/     Level: 1
Line:     main.go Level: 2
```
#### Hints:

```trimmed := strings.TrimLeft(line, " ")```

```indent := len(line) - len(trimmed)```

```level := indent / 2```

##### 💡 Task: Write the program that prints Line: + Level:.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

    trimmed := strings.TrimLeft(line, " ")
    indent := len(line) - len(trimmed)
    level := indent / 2

    fmt.Printf("Line: %s  Level: %d", strings.ToUpper(line), level)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
```

### 🔹 Mini Exercise (hands-on)

#### Run this program with a tree input like:
```bash
project/
  cmd/
    main.go
  internal/
    app.go
```
Verify that each line prints uppercase + correct level

#### Then write in you book:

1. What is the level of main.go?

2. What is the level of internal/?

## 🟢 Phase 3: Build the stack & path

**Next step:** we use a stack ([]string) to keep track of the path as we go deeper or shallower in the tree.

#### Remember:

* Each folder/file you read has a level

* If the level is less than the stack length, you pop the **stack** until it matches the level

* Then push the current name

* ```Full path = filepath.Join(stack...)```

### Exercise 3 (Hands-on)

#### Modify your code to:

Keep a var pathStack []string

**For each line:**

* Trim / if it’s a folder

* Adjust stack based on level

* Append current name to stack

#### Print the full path like:
```bash
Full path: project/cmd
Full path: project/cmd/main.go
Full path: project/internal
Full path: project/internal/app.go
```
#### Hints:

```strings.TrimSuffix(trimmed, "/")``` → removes trailing / for dirs

```filepath.Join(pathStack...)``` → build the full path

```If level < len(pathStack)```, slice the stack: ```pathStack = pathStack[:level]```

##### 💡 Task: Write your code to print full paths only (no file creation yet) and paste it here.

```go
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
	

		name := strings.TrimSuffix(trimmed, "/")

		
		if level < len(pathStack) {
			pathStack = pathStack[:level]
		}
		
		pathStack = append(pathStack, name)
		
		FullPath := filepath.Join(pathStack...)
		fmt.Println("Full path:", FullPath)
	}
	
	if err := scanner.Err(); err != nil{
		fmt.Println("Error reading input:", err)
	}
}
```

### 🔹 Mini Exercise (your turn)

#### Run the program with this input:
```bash
project/
  cmd/
    main.go
  internal/
    app.go
```

**Then answer:**

* What is the stack after reading cmd/?

* What is the stack after reading main.go?

* What is the full path printed for app.go?

This will check that you understand stack logic + path building.


## 🟢 Phase 4: Creating Directories and Files (Step-by-Step)

**Goal:**

Turn the full paths we calculated into real folders and files on disk.

### We’ll do it in two steps:

### ✅ Step 4a: Creating directories

**Concept:**

* If a line ends with / → it’s a folder

* Use ```os.MkdirAll``` to create the folder and any missing parents

* Permissions ``0755`` → owner can read/write/execute, others read/execute

#### Hints

* Use ``strings.HasSuffix(name, "/")`` to check if it’s a directory

* Use `filepath.Join(stack...)` to get the full path

* Handle errors with `if err != nil`

#### ✅ Phase 4a Exercises (No code — just reasoning)
##### Exercise 1

**You detect directories using:**

`strings.HasSuffix(trimmed, "/")`

##### Question:
Why must you check trimmed and not name?

##### Exercise 2

**Given the stack:**
```
pathStack = ["project", "internal"]
name = "handlers"
isDir = true
```
* What should the full path be?

* Write the answer exactly as a string.

##### Exercise 3

**Imagine:**
```
isDir == true
fullPath == "project/internal/handlers"
```
Which function will create the directory?

A) `os.Create(fullPath)`

B) `os.Mkdir(fullPath, 0755)`

C) `os.MkdirAll(fullPath, 0755)`

D) `filepath.Join(fullPath)`

Choose the correct letter.

##### Exercise 4

When making directories, why must we use MkdirAll instead of Mkdir?

Explain in 1–2 sentences.


### ✅ Phase 4b — Creating directories & files (Step-by-Step Learning)

Before writing code, we break the logic into tiny parts.

#### 🧠 Step 1 — What you must build

**Inside the loop, for every line:**
```go
If isDir == true

→ create a folder

If isDir == false

→ create a file
```
**We already know how to compute:**

`isDir`

`fullPath`

`pathStack`

Now you learn how to create things.

#### 🧠 Step 2 — How to create a directory (concept)
`os.MkdirAll(fullPath, 0755)`

**This means:**

* "Make this directory and any parents"

`0755` = permissions (owner read/write/execute, others read/execute)

**This works for:**

* project/internal/handlers

* even if internal/ didn't exist yet.

#### 🧠 Step 3 — How to create a file (concept)
`f, err := os.Create(fullPath)`

**This:**

* creates the file and overwrites it if it exists

* returns a file handle

* must be closed:

`defer f.Close()`

#### 🧠 Step 4 — Add this logic to your loop (pseudo-code)

**Not real Go code yet — just the mental blueprint:**
```go
if isDir:
    MkdirAll(fullPath)
else:
    Create(fullPath)
```
That's it.

You understand this?

#### 📝 Now your exercise (Phase 4b)

No code writing yet — just thinking.

##### Exercise 1 — Predict the output

**Input:**
```bash
myapp/
  cmd/
    main.go
  pkg/
    utils.go
```

Write EXACTLY what paths will be created
(one per line, in order).

##### Exercise 2 — Directory or file?

**For each, write DIR or FILE:**
```
backend/

backend/api/

backend/api/server.go

README.md

src/config/
```
##### Exercise 3 — Explain this in your own words

Why must the file creation happen after computing pathStack?

(1–2 sentences)

##### Exercise 4 — Error handling

What happens if you try `os.Create("src/config")`
when "src/config" is supposed to be a directory?

#### Code:

```go
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
			fmt.Println("\nProcessing tree...\n")
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
```

## 🟢 Phase 5: Export Existing Project Tree (--export mode)

**Goal:**

Allow the tool to read an existing folder on disk and output a tree-like structure that can be pasted into your Treemaker input.

This is the reverse of Phase 4: instead of creating files/folders, we scan and display them.

**Concept:**

**Use command-line flags to enable export mode:**

```go
var exportMode = flag.Bool("export", false, "export project tree instead of creating it")
flag.Parse()
```

**Running:**

```bash
go run main.go --export /path/to/project
```

sets `*exportMode = true`.

**Use `filepath.WalkDir` to recursively read the folder:**

```go
filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
    if err != nil {
        return err
    }

    // determine indentation
    // append "/" if it's a directory
    return nil
})
```

`path` → full path of the entry

`d.IsDir()` → tells you if it’s a directory

`err` → always handle errors

**Calculate depth for pretty printing:**

```go
depth := strings.Count(path, string(os.PathSeparator)) - baseDepth
prefix := strings.Repeat("  ", depth) // 2 spaces per level
```
**`Append` / for directories to match input format:**

```go
if d.IsDir() {
    fmt.Println(prefix + d.Name() + "/")
} else {
    fmt.Println(prefix + d.Name())
}
```

**Optional: write to a file:**

```go
file, _ := os.Create("tree.txt")
defer file.Close()
fmt.Fprintln(file, outputLine)
```

This allows the user to copy-paste this tree into Treemaker to recreate it elsewhere.

### Mini Exercise:

**Suppose your folder `/home/user/myapp` contains:**

```
myapp/
  cmd/
    main.go
  pkg/
    utils.go
  README.md
```

* What should `--export` output look like exactly, including indentation and / for directories?

* If a file is directly under the root, does it get / at the end? (Answer: No — only directories get /)

* How would you calculate depth for utils.go?