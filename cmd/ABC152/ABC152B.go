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
	a := s[0]
	na, _ := strconv.Atoi(a)
	b := s[1]
	nb, _ := strconv.Atoi(b)
	d := make([]string, 0)
	if na < nb {
		for i := 0; i < nb; i++ {
			d = append(d, a)
		}
		fmt.Fprintln(writer, strings.Join(d, ""))
	} else {
		for i := 0; i < na; i++ {
			d = append(d, b)
		}
		fmt.Fprintln(writer, strings.Join(d, ""))
	}
}
