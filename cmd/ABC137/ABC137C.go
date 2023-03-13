package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	ad := make([][]string, 100000)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		sc.Scan()
		s := strings.Split(sc.Text(), "")
		sort.Strings(s)
		ad[i] = s
	}
	var count int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if cmp(ad[i], ad[j]) {
				count += 1
			}
		}
	}
	fmt.Fprintln(writer, count)
}

func cmp(a, b []string) bool {
	for i := 0; i < 10; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
