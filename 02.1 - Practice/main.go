package main

import (
	"fmt"
)

func main() {
	fmt.Println("Practice")

	var tmp []int
	for i := 0; i <= 10; i++ {
		tmp = append(tmp, i)
	}
	fmt.Println(tmp)

	for _, value := range tmp {
		if value%2 == 0 {
			fmt.Println(fmt.Sprintf("%v is EVEN", value))
		} else {
			fmt.Println(fmt.Sprintf("%v is ODD", value))
		}
	}
}
