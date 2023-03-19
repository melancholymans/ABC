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
	n, _ := strconv.Atoi(s[0])
	k, _ := strconv.Atoi(s[1])
	for i := 0; i < k; i++ {
		if n%200 == 0 {
			n = n / 200
		} else {
			n = n*1000 + 200
		}
	}
	fmt.Fprintln(writer, n)
}
