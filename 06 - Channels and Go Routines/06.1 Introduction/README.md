# Website Status Checker

This is a simple Go program that checks the availability of a list of websites.

## Description

The program iterates through a predefined list of URLs and sends an HTTP GET request to each one. 

It then prints to the console whether the website is "up" (reachable) or "down" (unreachable or error).

This project is an introduction to Go routines and channels, as it emphasizes how this programs runs sequentially and can be a real issue if the list of websites is large.

## How to Run

1.  Ensure you have Go installed on your machine.
2.  Navigate to the project directory.
3.  Run the following command:

    ```bash
    go run main.go
    ```

## Output

The program will output the status of each website in the list, for example:

```
https://google.com  is up
https://facebook.com  is up
...
```

## Code Structure

-   `main.go`: Contains the main logic for defining the links and checking their status.
    -   `main()`: Entry point of the application.
    -   `checkLink(link string)`: Helper function to perform the HTTP GET request and print the status.
