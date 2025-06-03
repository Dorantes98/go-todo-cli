# Go Notes
## package main
- This tells Go: "This file belongs to the main program."
- It's required for any file that’s going to produce a runnable binary (CLI or app).

## Improrts
- `import` is used to bring in a package or module so you can use its functions and variables
- `import` can be used to import a specific function or variable from a package

### 'fmt' package
- contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.

### 'time' package
- provides functionality for measuring and displaying time. It includes functions for getting the current time, formatting dates, and performing time calculations.

### 'os' package
- provides functions for interacting with the operating system, such as reading and writing files, handling command-line arguments, and managing environment variables.

## Code
```go
func main()
```
- The entry point of a Go program. When you run the program, this function is executed first. Ex. _go run main.go_

```go
fmt.Println("Todo List:")
```
- Standard Go function to print text to stdout.
- Simpple test that the cli is working.

```go
for i, arg := range os.Args {
  fmt.Printf("%d: %s\n", i, arg)
}
```
- This for loop iterates over the command-line arguments passed to the program.
- `os.Args` is a slice of strings that contains the command-line arguments.
- `i` is the index of the argument, and `arg` is the value of the argument.

```go
if os.Args[1] == "add" {}
```
- This checks if the first command-line argument is "add". If it is, the code inside the curly braces will be executed.

```go
file, err := os.OpenFile("tasks.csv", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
```
- This line opens a file named "tasks.csv" for appending data. If the file does not exist, it will be created. The `0644` sets the file permissions.
  - `0644` means:
    - Owner can read and write (6)
    - Group can read (4)
    - Others can read (4)
- `os.OpenFile` returns a file pointer and an error. If the file is opened successfully, `err` will be `nil`.
- If there is an error opening the file, `err` will contain the error information.
- `os.O_APPEND` means that data will be written at the end of the file.
- `os.O_CREATE` means that if the file does not exist, it will be created. 
- `os.O_WRONLY` means that the file will be opened for writing only.
- `os.O_RDWR` would mean the file is opened for reading and writing, but in this case, we only want to write to it.

# Commit messages
## Common Prefix Types:
- Prefix:	    Meaning
- feat:	      A new feature
- fix:	      A bug fix
- docs:	      Documentation only changes
- style:	    Code style changes (formatting, etc)
- refactor:	  Code changes that neither fix a bug nor add a feature
- perf:	      Performance improvements
- test:	      Adding or updating tests
- build:	    Changes that affect the build system or dependencies
- ci:	        Continuous integration-related changes
- chore:	    Other changes that don’t modify src or test files (e.g., release tasks)
- revert:	    Revert a previous commit