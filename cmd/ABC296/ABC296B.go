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
	m := make(map[int]string)
	m[0] = "a"
	m[1] = "b"
	m[2] = "c"
	m[3] = "d"
	m[4] = "e"
	m[5] = "f"
	m[6] = "g"
	m[7] = "h"
	for i := 0; i < 8; i++ {
		sc.Scan()
		s := strings.Split(sc.Text(), "")
		for j := 0; j < 8; j++ {
			if s[j] == "*" {
				fmt.Fprintln(writer, m[j]+strconv.Itoa(8-i))
			}
		}
	}
}
