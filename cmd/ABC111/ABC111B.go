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
	lst := []int{111, 222, 333, 444, 555, 666, 777, 888, 999}
	for _, v := range lst {
		if n <= v {
			fmt.Fprintln(writer, v)
			return
		}
	}
}
