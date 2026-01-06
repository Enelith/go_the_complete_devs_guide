# Assignment - Interfaces

This project is an assignment focused on understanding and implementing Interfaces in Go.

## Getting Started

- Write a program that creates two custom struct types called `triangle` and `square` 
- The `square` type should be a struct with a field called `sideLength` of type `float64`
- The `triangle` type should be a struct with fields called `base` and `height` of type `float64`
- Both types should have a function called `getArea` that returns calculated area of the square and the triangle.
  - Area of a square is `sideLength * sideLength`
  - Area of a triangle is `base * height / 2` or `base * height * 0.5`
- Add a `shape` interface that defines a function called `getArea`. 
  - Design the interface so that the `getArea` function can be called with either a `square` or a `triangle`. 
- Create a function `printArea`. 
  - This function should calculate the area of the given shape and print it out to the terminal.


### Prerequisites

- Go installed on your machine.

### Running the program

```bash
go run main.go
```
