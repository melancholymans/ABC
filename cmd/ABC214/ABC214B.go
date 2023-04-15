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
	s, _ := strconv.Atoi(r[0])
	t, _ := strconv.Atoi(r[1])
	count := 0
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			for c := 0; c <= 100; c++ {
				if a+b+c <= s && a*b*c <= t {
					count += 1
				}
			}
		}
	}
	fmt.Fprintln(writer, count)
}
