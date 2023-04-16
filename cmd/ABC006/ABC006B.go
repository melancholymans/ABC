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
	r := sc.Text()
	n, _ := strconv.Atoi(r)
	f0 := 0
	f1 := 0
	f2 := 1
	if n-1 == 0 || n-1 == 1 {
		fmt.Fprintln(writer, 0)
		return
	}
	if n-1 == 2 {
		fmt.Fprintln(writer, 1)
		return
	}
	for f := 3; f < n; f++ {
		f3 := (f0 + f1 + f2) % 10007
		f0 = f1
		f1 = f2
		f2 = f3
	}
	fmt.Fprintln(writer, f2)
}
