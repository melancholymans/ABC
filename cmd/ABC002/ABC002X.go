package main

import (
	"fmt"
)

func main() {
	sanples := []string{"hello", "apple", "strings"}
loopouter:
	for _, sample := range sanples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' || r == 'e' {
				break loopouter
			}
		}
		fmt.Println()
	}
	fmt.Println("抜けた後はここに来る")
}
