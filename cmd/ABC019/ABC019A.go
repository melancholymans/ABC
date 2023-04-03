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
	ad := make([]int, 0)
	for i := 0; i < 3; i++ {
		a, _ := strconv.Atoi(s[i])
		ad = append(ad, a)
	}
	sort.Sort(sort.IntSlice(ad))
	fmt.Fprintln(writer, ad[1])
}
