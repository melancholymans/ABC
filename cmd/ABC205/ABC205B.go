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
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	a := strings.Split(sc.Text(), " ")
	ad := make([]int, 0)
	for i := 0; i < n; i++ {
		m, _ := strconv.Atoi(a[i])
		ad = append(ad, m)
	}
	sort.Sort(sort.IntSlice(ad))
	answer := make([]int, 0)
	for i := 1; i <= n; i++ {
		answer = append(answer, i)
	}
	for i := 0; i < n; i++ {
		if answer[i] != ad[i] {
			fmt.Fprintln(writer, "No")
			goto Z
		}
	}
	fmt.Fprintln(writer, "Yes")
Z:
}
