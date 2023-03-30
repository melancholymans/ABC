package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	ad := make([]int, 0)
	m := make(map[int]bool)
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		if !m[a] {
			m[a] = true
			ad = append(ad, a)
		}
	}
	sort.Sort(sort.IntSlice(ad))
	fmt.Fprintln(writer, ad[len(ad)-2])
}
