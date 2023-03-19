package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	s := strings.Split(sc.Text(), " ")
	a, _ := strconv.Atoi(s[0])
	b, _ := strconv.Atoi(s[1])
	c, _ := strconv.Atoi(s[2])
	d, _ := strconv.Atoi(s[3])
	for i := 0; ; i++ {
		c = c - b
		if c <= 0 {
			fmt.Fprintln(writer, "Yes")
			break
		}
		a = a - d
		if a <= 0 {
			fmt.Fprintln(writer, "No")
			break
		}
	}
}
