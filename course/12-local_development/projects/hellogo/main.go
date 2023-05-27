package main

import (
	"fmt"

	"github.com/hwangblood/mystrings"
)

func main() {
	fmt.Println("hello world")

	reversedStr := mystrings.Reverse("hello world")
	fmt.Println(reversedStr)
}
