# Website Status Checker

This is a simple Go program that demonstrates the use of **Goroutines** and **Channels** to check the status of multiple websites concurrently.

## Description

The program iterates through a list of website URLs and spawns a new Goroutine for each URL to check if it is up or down. It uses a Channel to communicate the result back to the main routine.

## Features

- **Concurrency**: Checks multiple websites at the same time using Goroutines.
- **Channels**: Demonstrates how to send messages between Goroutines.
- **HTTP Requests**: Performs simple HTTP GET requests to check site availability.

## Theory of Go Routines:

How to create a goroutine?
```
go checkLink(link)
```
- `go`: Create a new thread go routine... (this `go` keyword is a built-in function and can only be placed before function calls)
- `checkLink(link)`: ...and run the function in the new thread.


Behind the scenes, there's what we call the Go Scheduler. It's responsible for deciding which goroutine to run next.
It's also responsible for managing the goroutines lifecycle.

By default, in Go, the scheduler runs all goroutines concurrently, and 1 goroutine uses 1 CPU core.

- ***Concurrency***:
  <br/>We can have multiple threads executing code. If one thread blocks, another one is picked up and worked on.
  We're scheduling which thread is going to be executed next.

- ***Parallelism***:
  <br/>Multiple threads executed at the same exact time, and working on the same data at the same time. Requires multiple CPU cores.
  We litteraly have multiple threads executing at the same time.

### Main Routines VS Child Routines
| Our Running Program |                                                     |
|:--------------------|:----------------------------------------------------|
| `Main Routine`      | `Main Routine created when we launched our program` |
| `Child Routine`     |                                                     |
| `Child Routine`     | `Child Routines created with the 'go' keyword`      |
| `Child Routine`     |                                                     |

By default, whenever the Main Routine finishes, all Child Routines are killed, and it will feel like the program stopped or didn't worked out.
<br/>
To prevent this, we need to use another construct in Go, called: **Channels**.

## Channels
Channels are used to communicate between different running goroutines, and it's the only way to do so.

Channels will be used to notify the Main Routine when all Child Routines have finished their work (and prevent the program from exiting before all Child Routines have finished).

Note: Channels are **typed**.
<br/>The data shared between goroutines is called ***shared memory***, and must be of the same type.
<br/>The data type of a channel is `chan T`, where `T` is the type of the data shared between goroutines.

### Sending Data with Channels
| Syntaxe Examples          | Description                                                                                     |
|:--------------------------|:------------------------------------------------------------------------------------------------|
| `channel <- 5`            | `Send the value '5' into the channel`                                                           |
| `myNumber <- channel`     | `Wait for a value to be sent into the channel. When we get one, assign the value to 'myNumber'` |
| `fmt.Println(<- channel)` | `Wait for a value to be sent into the channel. When we get one, log it out immediately`        |

## How to Run

1.  Ensure you have Go installed on your machine.
2.  Navigate to the project directory.
3.  Run the following command:

    ```bash
    go run main.go
    ```

## Code Overview

-   `main()`: Initializes a list of links and a channel. It launches a `checkLink` goroutine for each link.
-   `checkLink(link string, c chan string)`: Performs an HTTP GET request to the specified link. It sends a message to the channel indicating whether the site is up or down.

## Note

This example is part of a learning exercise on Go concurrency patterns.
