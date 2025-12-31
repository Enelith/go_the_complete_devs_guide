# Structs with Pointers in Go

This project demonstrates how to work with structs and pointers in Go, specifically focusing on updating struct fields using pointer receivers. It is part of the "Go: The Complete Developer's Guide" course by Stephen Grider.

## Key Concepts

*   **Struct Definition**: Defining custom types using `struct`.
*   **Embedded Structs**: Embedding one struct (`contactInfo`) inside another (`person`).
*   **Pointer Receivers**: Using pointers (`*person`) in function receivers to allow modifying the underlying value.
*   **Value vs Pointer**: Understanding the difference between passing by value and passing by reference (pointer).
*   **Generic Printing**: A helper function demonstrating how to accept any type using `interface{}` and inspecting runtime information.

## Value vs Pointer
- &variable: A pointer to a variable. Give me the memory address of the value this variable is pointing at.
- *variable: A pointer to a pointer. Give me the value this memory address is pointing at.
## Code Overview

The `main.go` file contains:

*   `person` and `contactInfo` struct definitions.
*   `updateName`: A method with a pointer receiver (`*person`) that updates the `firstName`.
*   `print`: A method with a value receiver to display the struct.
*   `genericPrint`: A function that takes an `interface{}` to print details about any passed variable, along with the calling function name.

## Usage

To run the program:

```bash
go run main.go
```

## Example Output

```text
main.genericPrint: {firstName:Jim lastName:Halpert contactInfo:{email:test@test.com zipcode:95128}}
{firstName:Jimmy lastName:Halpert contactInfo:{email:test@test.com zipcode:95128}}
```
