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
	st := []byte(sc.Text())
	limit := len(st) / 2
	for i := 0; i < limit; i++ {
		a := st[2*i]
		b := st[2*i+1]
		st[2*i] = b
		st[2*i+1] = a
	}
	fmt.Println(string(st))
}
