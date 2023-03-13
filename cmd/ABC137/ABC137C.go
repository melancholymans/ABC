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
	ad := make([]int, 100000)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		total := 0
		sc.Scan()
		s := []byte(sc.Text())
		for j := 0; j < 10; j++ {
			total += int(s[j])
		}
		ad[i] = total
	}
	var count int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if ad[i] == ad[j] {
				count += 1
			}
		}
	}
	fmt.Fprintln(writer, count)
}

/*
func cmp(a, b []string) bool {
	for i := 0; i < 10; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
*/
