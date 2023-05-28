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
	fmt.Fprintln(writer, numberDigits(n))
}

func numberDigits(n int) int {
	for i := 1; ; i++ {
		n = n / 10
		if n == 0 {
			return i
		}
	}
}
