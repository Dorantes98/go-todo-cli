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