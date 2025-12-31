# Go Structs Example

This project demonstrates the basics of defining and using structs in Go.

## Overview

The `main.go` file contains a simple program that defines a `person` struct and illustrates three different ways to create and initialize it.

## Code Highlights

### Struct Definition

```go
type person struct {
    firstName string
    lastName  string
}
```

### Initialization Methods

1.  **Positional Arguments**: Initializing a struct by passing values in the order of fields.
    ```go
    alex_1 := person{"Alex", "Smith"}
    ```

2.  **Named Arguments**: Initializing a struct by specifying field names (preferred for clarity).
    ```go
    alex_2 := person{firstName: "Alex", lastName: "Smith"}
    ```

3.  **Zero Values**: Declaring a struct variable without initialization sets fields to their zero values (e.g., empty string `""` for strings).
    ```go
    var alex_3 person
    ```

### Updating Fields

Fields can be accessed and updated using the dot notation:

```go
alex_3.firstName = "Alexander"
alex_3.lastName = "Smith"
```

## Running the Code

To run the program, execute the following command in your terminal:

```bash
go run main.go
```
