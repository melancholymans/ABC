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
	n, _ := strconv.Atoi(sc.Text())
	m := map[string]int{}
	sl := []string{}
	for i := 0; i < n; i++ {
		sc.Scan()
		w := sc.Text()
		m[w] += 1
		sl = append(sl, w)
	}
	for _, v := range m {
		if v > 1 {
			fmt.Fprintln(writer, "No")
			return
		}
	}
	for i := 0; i < n-1; i++ {
		if sl[i][len(sl[i])-1] != sl[i+1][0] {
			fmt.Fprintln(writer, "No")
			return
		}
	}
	fmt.Fprintln(writer, "Yes")
}
