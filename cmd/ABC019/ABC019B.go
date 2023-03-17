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
	s := sc.Text()
	i, j := 0, 0
	n := len(s)
	var result string
	for {
		c := s[j]
		if i == j {
			result = result + string(c)
		}
		j += 1
		if j >= n {
			result = result + strconv.Itoa(j-i)
			break
		}
		if s[i] != s[j] {
			result = result + strconv.Itoa(j-i)
			i = j
		}
	}
	fmt.Fprintln(writer, result)
}
