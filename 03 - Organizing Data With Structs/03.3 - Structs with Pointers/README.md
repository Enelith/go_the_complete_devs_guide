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

## Reference and Value Types

In the following example,
```bash
func main() {
  tmp := []string{"One", "Two", "Three"}
  
  modify(tmp)
  
  fmt.Println(tmp)
}

func modify(s []string) {
  s[0] = "Four"
}
```
The result of fmt.Println will be `[]string{"Four","Two","Three"}` and not `[]string{"One","Two","Three"}`, because the value of s is copied into modify function.

It is important to understand that `[]string` is a slice, but it's also a *reference type*.

As a reminder:
- An Array is a Primitive data structure, it can't be resized, and rarely used directly.
- A Slice can grow and shrink dynamically. It is used  99% of the time for lists of elements.

The Slice (as a Reference type), when created, generate in the background 2 structures in 2 different memory address (an array and a slice).
- The slice structure roughly contains the following: a length (number of elements), a capacity (maximum number of elements), and a pointer to the array.
- The actual array containing all the elements the Slice is referencing.

So, when passing the Slice as an argument to the function, the Slice structure is copied into the function, but the reference to the array stays the same.
Therefore, the change made is being reflected in the original Slice.


| Value Types | Reference Types |
| :--- | :--- |
| `int`, `float`, `string`, `bool` | `slice` |
| `struct` | `map` |
| `array` | `channel` |
| | `pointer` |
| | `function` |

- **Value Types**: You need to use Pointers to change these things in a function.
- **Reference Types**: You don't need to worry about pointers with these.

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
