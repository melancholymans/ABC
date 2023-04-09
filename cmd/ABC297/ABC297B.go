package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	r := sc.Text()
	//m := make(map[string]int)
	for i, v := range r {
		m[string(v)] = i
	}
	fmt.Println(m)
	/*
		if m["K"] == 1 && m["Q"] == 1 && m["R"] == 2 && m["B"] == 2 && m["N"] == 2 {

			fmt.Println("OK")
		} else {
			fmt.Println("No")
		}
	*/
}
