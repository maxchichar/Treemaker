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

## 🟢 Phase 5 — Support Pretty ASCII Tree Formats

Your treemaker currently understands this:
```
myapp/
  cmd/
    main.go
```
But it cannot understand this:
```
myapp/
├── cmd/
│   └── main.go
└── internal/
    └── service.go
```
This phase teaches you to support both formats using a **normalization layer.**

### Part A — Understanding The Problem (Concept Exercises)

#### Exercise 1 — Identify all ASCII tree symbols

Write down (on paper or inside your Guide.md) all the characters that appear in pretty trees:

Example:
```
├──
└──
│
```
❓ **Question:**
List at least 4 ASCII tree symbols your treemaker must remove or transform.

#### Exercise 2 — Compare indentation rules

Look at these two lines:
```
  cmd/
│   └── main.go
```
**❓ Questions:**

*1.* Which one uses **spaces** for indentation?
*2.* Which one uses **ASCII tree glyphs** to represent indentation?
*3.* Why must both be converted into the same indentation format before your depth logic works?

Write 2–3 sentences.

#### Exercise 3 — Plan a normalization function

Before writing code, outline in plain English:

* What should the normalization step do?
* What must it remove?
* What must it keep?
* Why must this run before indentation detection?

Write 3–5 sentences.

### Part B — Hands-On Input Transformation

You will now manually transform ASCII tree lines into the simple format your program understands.

#### Exercise 4 — Manual conversion

Transform each line manually.

Input:
```
├── src/
│   ├── main.go
│   └── utils.go
└── README.md
```
Write manually what each line should look like after normalization (spaces only, no ASCII symbols).

Expected pattern:
```
src/
  main.go
  utils.go
README.md
```
You must produce the exact indentation yourself.

#### Exercise 5 — Count spaces

For each of the normalized lines you wrote above, count:

* number of left spaces
* calculated level (spaces ÷ 2)

Example answer format:
```
"  main.go" → 2 spaces → level 1
```

#### Exercise 6 — Detect your algorithm

Based on your own transformations, answer:

❓ What are the exact steps to convert an ASCII tree line to your standard format?

Write down the algorithm in bullet points.

Example style (not the solution):

* Remove X
* Convert Y to Z
* Trim W
* Preserve indentation

### Part C — Build the Normalizer in Go (step-by-step exercises)

#### Exercise 7 — Write a function skeleton

Write a function signature only:
```
func normalizeLine(line string) string
```
Do NOT implement it yet.

#### Exercise 8 — Replace only one symbol

Modify normalizeLine so it removes just one symbol:
```
│
```
Test input:
```
│   └── main.go
```
Expected output (after your manual reasoning — NOT code from me):
```
   └── main.go
```
Run your program printing normalized lines only.

#### Exercise 9 — Add second symbol

Now modify normalizeLine to remove:
```
├──
```
Test input:
```
├── src/
```
Expected output:
```
src/
```
Verify your logic.

#### Exercise 10 — Add the remaining symbols

Make your function handle all ASCII symbols you identified in Exercise 1.

Test with this tree:
```
app/
├── cmd/
│   ├── main.go
│   └── helper.go
└── README.md
```
Print ONLY the normalized lines (not paths yet).

#### Exercise 11 — Confirm indentation levels

Modify your program temporarily to print:
```
[normalized] -> [spaces] -> [level]
```
Example:
```
"  main.go" -> 2 spaces -> level 1
```
Make sure the levels match the logical depth of the ASCII tree.

If not → fix normalization.

### Part D — Integrating With Treemaker

#### Exercise 12 — Full pipeline test

Paste a complex pretty tree:
```
backend/
├── api/
│   ├── server.go
│   └── router.go
├── internal/
│   └── utils.go
└── README.md
```
Verify the following:

* Normalized output (spaces only)
* Correct indent level detection
* Correct pathStack behavior
* Correct directory/file creation

Fix anything that breaks.

#### Exercise 13 — Mixed tree modes

Paste this hybrid tree:
```
backend/
  api/
    routes.go
├── pkg/
│   └── helper.go
```
Your treemaker must:

* support mixed indentation
* still create correct paths

Debug until it works.

### Part E — Final Challenge

#### Exercise 14 — Write your own rules

Write:
```
THE 5 RULES OF NORMALIZING ASCII TREES
```
These should be your own distilled principles after building and debugging the feature.

#### Exercise 15 — Build the final implementation

Now and ONLY now implement your final normalizeLine() function using all the knowledge you gathered.

Run a full project tree through it and verify the output.

#### Exercise 16 — Document Phase 5 in Guide.md

Add the following sections:

* What ASCII trees are
* Why normalization is required
* How your algorithm works (in your own words)
* Example conversions
* Exercises
* Your final code (only after completing exercises)