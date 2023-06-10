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
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	r1 := strings.Split(sc.Text(), " ")
	count := 0
	odd_count := 0
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(r1[i])
		if a%2 == 0 {
			odd_count += 1
			if a%3 == 0 || a%5 == 0 {
				count += 1
			}
		}
	}
	if count == odd_count {
		fmt.Fprintln(writer, "APPROVED")
	} else {
		fmt.Fprintln(writer, "DENIED")
	}
}
