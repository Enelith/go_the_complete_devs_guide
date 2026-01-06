# Hard Mode Interfaces Assignment

This project is a solution for the "Hard Mode Interfaces" assignment from Stephen Grider's Go course.

## Getting Started

- Create a program that reads the contents of a text file then prints its content to the terminal.
- The file to open should be provided as an argument to the program when it is executed at the terminal (e.g. `go run main.go myfile.txt` should open the myfile.txt file).
- To read in the arguments provided to a program, you can reference the variable `os.Args` in Go, which is a slice of type string.
- To open the file, check out the documentation for the `os.Open` function (https://pkg.go.dev/os#Open)
- What interface does the `File` type implement? Hint: os.File implements io.Reader.
- If the `File` type implements the `io.Reader` interface, you might be able to reuse the `io.Copy` function to read the contents of the file into a string. 

### Prerequisites

- Go installed on your machine.

### Running the program

```bash
go run main.go
```
