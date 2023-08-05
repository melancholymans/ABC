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
	r := strings.Split(sc.Text(), " ")
	a := r[0]
	b := r[1]
	c, _ := strconv.Atoi(a + b)
	for i := 1; i < c; i++ {
		if i*i == c {
			fmt.Fprintln(writer, "Yes")
			return
		}
	}
	fmt.Fprintln(writer, "No")
}
