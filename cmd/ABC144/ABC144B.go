package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	calc(n)
}

func calc(n int) {
	for x := 1; x <= 9; x++ {
		for y := 1; y <= 9; y++ {
			if n == x*y {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}
