# A Simple Start

This is a simple Go program that prints "Hello World" to the console. It serves as a basic introduction to Go programming.

## How to Run

To run this program, execute the following command in your terminal:

```bash
go run main.go
```

## Notions
### 1) What does "package main" mean ?
There are 2 kind of packages in Go: 
- packages that contain code that can be imported by other packages (REUSABLE packages)
- packages that contain code that can be executed directly (main packages) (EXECUTABLE packages)

The main package is the one that is executed when you run a Go program.
It defines a package that can be compiled and then "executed". **Must have a func called 'main'**

A REUSABLE package can be imported by other packages. It defines a package that can be used as a dependency (helper code).
(example: package calculator, package uploader, ...)

### 2) What does "import 'fmt'" mean ?
It imports the fmt package, aka give my package main all the functions that are in the fmt package.

fmt is a package that contains functions to print to the console. It's a standard library package, that is included with every Go installation.

Many other packages are available in the standard library, including packages for working with files, networking, ...
Standard official documentation can be found here: https://golang.org/pkg/

### 3) What's that 'func' thing ?
It's a function declaration. 'func' is short for 'function'.

A simple structure for a func is as follow: 
func functionName(parameters) returnType {
	// code
}

### 4) How is the main.go file organized ?
At the very top, you'll find the package declaration: for example, package main.

Right underneath that, you'll find the import declaration needed for your file, in this case, import "fmt" (but can also include reusable packages).

We then have the body of the file, which is where we add a bunch of logic, that actually kind of does something.
So it will be a collection of different functions, statements, variables declarations, ...

So in general, we're going to be very used to this very same pattern of code, in every Go file. 

## References
- https://tour.golang.org/welcome/1

