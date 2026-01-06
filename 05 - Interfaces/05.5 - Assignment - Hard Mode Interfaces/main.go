package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Args[1])

	fileName := os.Args[1]

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	// Note that os.File implements io.Reader
	_, _ = io.Copy(os.Stdout, f)
}
