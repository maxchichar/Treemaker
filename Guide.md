# GUIDE ON HOW TO BUILD A TREEMAKER

---
### ***⚠️🛑 Complete The Exercise Before Checking The Correct Code Below The Exercise.***
---

## 🟢 Phase 1: Read input line by line

#### Goal: Learn to read a multi-line tree from stdin.

### Exercise 1:

#### Write a Go program that:

Reads lines from standard input (os.Stdin)

Prints each line prefixed with "Line: "

#### Hints:

Use bufio.NewScanner(os.Stdin)

Use a for scanner.Scan() loop

Use scanner.Text() to get each line

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
use strings.ToUpper(line) (Study the strings package).

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

trimmed := strings.TrimLeft(line, " ")

indent := len(line) - len(trimmed)

level := indent / 2

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

Next step: we use a stack ([]string) to keep track of the path as we go deeper or shallower in the tree.

#### Remember:

* Each folder/file you read has a level

* If the level is less than the stack length, you pop the **stack** until it matches the level

* Then push the current name

* Full path = filepath.Join(stack...)

### Exercise 3 (Hands-on)

#### Modify your code to:

Keep a var pathStack []string

For each line:

Trim / if it’s a folder

Adjust stack based on level

Append current name to stack

#### Print the full path like:
```bash
Full path: project/cmd
Full path: project/cmd/main.go
Full path: project/internal
Full path: project/internal/app.go
```
#### Hints:

strings.TrimSuffix(trimmed, "/") → removes trailing / for dirs

filepath.Join(pathStack...) → build the full path

If level < len(pathStack), slice the stack: pathStack = pathStack[:level]

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

Then answer:

What is the stack after reading cmd/?

What is the stack after reading main.go?

What is the full path printed for app.go?

This will check that you understand stack logic + path building.