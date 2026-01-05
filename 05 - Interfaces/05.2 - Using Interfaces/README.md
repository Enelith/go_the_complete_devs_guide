# Go Interfaces Example

This project demonstrates the usage of interfaces in Go. 

## 5.60 - interfaces in Practice
To whom it may concern...
```
type bot interface
```
Our program has a new type called `bot`.

```
getGreeting() string
```
If you are a type in this program with a function called `getGreeting` and you return a string, then you are now an honorary member of type `bot`.

Now that you're also an honorary member of type `bot`, you can now call this function called `printGreeting` on any value of type `bot`.
```
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
```

## How to Run

To run the program, execute the following command in your terminal:

```bash
go run main.go
```

## Output

```text
Hello!
Hola!
```
