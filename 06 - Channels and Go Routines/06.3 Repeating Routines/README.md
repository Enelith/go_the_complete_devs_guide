# Website Status Checker (Repeating Routines)

This is a Go program that concurrently checks the status of a list of websites to see if they are up or down. It demonstrates the use of **Goroutines** and **Channels** to handle concurrency and repeat the checks indefinitely.

## Description

The program iterates through a slice of URLs and spawns a new goroutine for each URL to check its status via an HTTP GET request. 

Once a check is complete, the URL is sent back through a channel. The main function listens to this channel, and upon receiving a URL (indicating a finished check), it waits for 2 seconds before spawning a new goroutine to check that same URL again. This creates an infinite loop of status checks for each website.

## Key Concepts

- **Goroutines**: Used to perform HTTP requests concurrently without blocking the main execution flow.
- **Channels**: Used for communication between the main routine and the worker goroutines.
- **Function Literals (Anonymous Functions)**: Used to implement the delay logic. By wrapping the sleep and the recursive call in a separate goroutine, the main loop remains unblocked.
- **Infinite Loop**: The program uses a `for range` loop over the channel to continuously process completed checks.

## Websites Monitored

The following websites are checked:
- https://google.com
- https://facebook.com
- https://humblebundle.com
- https://stackoverflow.com
- https://golang.org
- https://amazon.com

## How to Run

1. Ensure you have Go installed on your machine.
2. Navigate to the directory containing `main.go`.
3. Run the following command:

```bash
go run main.go
```

## Output

The program will print whether each link is "up" or "might be down". It will continue to do so every few seconds until you terminate the program (e.g., using `Ctrl+C`).

```text
Looping forever (and repeating each websites check after they've send their answer through the channel...)
https://google.com  is up!
https://stackoverflow.com  is up!
...
```
