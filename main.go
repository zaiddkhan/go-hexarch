package main

import (
	"fmt"
	"go-arch/internal/adapters/core/arithmetic"
)

func main() {
	arithAdapter := arithmetic.NewAdapter()
	result, err := arithAdapter.Addition(1, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
