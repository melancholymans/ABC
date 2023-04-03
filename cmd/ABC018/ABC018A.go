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
	ad := []int{}
	sd := []int{}
	for i := 0; i < 3; i++ {
		sc.Scan()
		n, _ := strconv.Atoi(sc.Text())
		ad = append(ad, n)
		sd = append(sd, n)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sd)))
	for _, v := range ad {
		fmt.Fprintln(writer, search(sd, v)+1)
	}
}

func search(a []int, t int) int {
	for i := 0; i < len(a); i++ {
		if a[i] == t {
			return i
		}
	}
	return -1
}
