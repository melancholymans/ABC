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
	in := sc.Text()
	i, j := 0, 0
	n := len(in)
	var result []byte
	for {
		o := in[j]
		if i == j {
			result = append(result, o)
		}
		j += 1
		if j <= n {
			l := strconv.Itoa(j - i)
			result = append(result, byte(l))
			i = j
		}
		if in[i] != in[j] {
			e := strconv.Itoa(j - i)
			append(result, e)
			i = j
		}
	}
	fmt.Fprintln(writer, in)
}
