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
	s := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(s[0])
	k, _ := strconv.Atoi(s[1])
	sc.Scan()
	a := strings.Split(sc.Text(), " ")
	ad := make([]int, n)
	for i := 0; i < n; i++ {
		ad[i], _ = strconv.Atoi(a[i])
	}
	sort.Sort(sort.Reverse((sort.IntSlice(ad))))
	total := 0
	for i := 0; i < k; i++ {
		total += ad[i]
	}
	fmt.Fprintln(writer, total)
}
