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
	a := strings.Split(sc.Text(), " ")
	ad := make([]int, n)
	for i := 0; i < n; i++ {
		ad[i], _ = strconv.Atoi(a[i])
	}
	count := 0
	for i := 1; i < n; i++ {
		crh := ad[i]
		if calc(ad, crh, i) {
			count += 1
		}
	}
	fmt.Fprintln(writer, count+1)
}

func calc(ad []int, crh, i int) bool {
	for j := 0; j < i; j++ {
		if ad[j] > crh {
			return false
		}
	}
	return true
}
