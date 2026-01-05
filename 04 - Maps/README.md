# Go Maps Example

This project demonstrates the usage of Maps in Go, based on the "Go: The Complete Developer's Guide" course by Stephen Grider.

## Features

The `main.go` file covers the following concepts:

1.  **Declaration**: Different ways to declare and initialize a map.
    *   Using `var` keyword.
    *   Using `make` function.
    *   Using map literals.
2.  **Manipulation**:
    *   Adding key-value pairs.
    *   Deleting key-value pairs using the `delete` function.
3.  **Iteration**:
    *   Iterating over a map using a `for` loop with `range`.

## Difference between a Map and a Struct

| Map                                                   | Struct                                                          |
|:------------------------------------------------------|:----------------------------------------------------------------|
| `All keys must be the same type`                      |                                                                 |
| `All values must be the same type`                    | `Values can be of different type`                               |
| `Use to represent a collection of related properties` | `Use to represent a "thing" with a lot of different properties` |
| `Don't need to know all the keys at compile time`     | `You need to  know all the different fields at compile time`    |
| `Keys are indexed - we can iterate over them`         | `Keys don't support indexing`                                   |
| `Reference Type !`                                    | `Value Type !`                                                  |

## How to Run

Ensure you have Go installed on your machine.

1.  Navigate to the project directory.
2.  Run the following command:

```bash
go run main.go
```

## Output

The program will print the state of the maps at various stages and finally iterate through the map to print each key-value pair.
