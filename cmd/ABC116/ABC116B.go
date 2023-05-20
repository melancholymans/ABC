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
	s, _ := strconv.Atoi(sc.Text())
	n := 1000000
	mp := map[int]bool{}
	mp[s] = true
	for i := 2; i <= n; i++ {
		s = calc(s)
		if mp[s] {
			fmt.Fprintln(writer, i)
			return
		}
		mp[s] = true
	}
}

func calc(s int) int {
	if s%2 == 0 {
		return s / 2
	} else {
		return 3*s + 1
	}
}
