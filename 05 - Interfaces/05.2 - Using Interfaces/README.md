# Go Interfaces Example

This project demonstrates the usage of interfaces in Go. 

## 5.60 - Interfaces in Practice
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

## 5.61 Rules of Interfaces
To reiterate, in our program: 
- We have a type called `englishBot`
  - `englishBot` has a function called `getGreeting` associated with it.

```
type englishBot struct{}
func (englishBot) getGreeting() string
```
When we say "associated", we mean the `getGreeting` function expects a type `englishBot` as its receiver (therefore, we say it is associated with the type `englishBot`).

- Likewise, we have same for the `spanishBot`
  - `spanishBot` has a function called `getGreeting` associated with it.
```
type spanishBot struct{}
func (spanishBot) getGreeting() string
```

We then defined a new interface called `bot`, which expects to see any other type inside of the application that implements a function called `getGreeting` that returns a string.
```
type bot interface{
    getGreeting() string
}
```

If any other type in the application DO implements this function, then that type is also considered of type `bot`.

For example, 
if our interface `bot` was as follow:
```
type user struct {
    name string
}

type bot interface {
    getGreeting(string, int) (string, error)
    getBotVersion() float64
    respondToUser(user) string
}
```
In order for any type to be considered of type `bot` as well, those types would need to implement all three functions (with same kind of arguments, and same returns).

### Terminology
- `Concrete Type`: it's something we can create a value out of directly (instanciate), and then access it, change it, or create new copies of it.
- `Interface Type`: Can't be created directly.

| Concrete Type                                     | Interface Type |
|:--------------------------------------------------|:---------------|
| `map` `struct`                                    | `bot`          |
| `int` `string`                                    |                |
| `englishBot`                                      |                |

### Extra Interface Notes
- `Interfaces are NOT generic types`: *Other languages have "generic" types (such as C#, Java...) - Go (famously) does not.*
- `Interfaces are 'implicit'`: *We don't manually have to say that our custom type satisfies some interface. (Can be a boon or a curse)*
- `Interfaces are a contract that help us manage types`: *GARBAGE IN -> GARBAGE OUT. If our custom type's implementation of a function is broken then interfaces won't help us!*
- `Interfaces are tough. Step#1 is understanding how to read them`: *Understand how to read interfaces in the standard library. Writing your own interfaces is tough and requires experience.*

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
