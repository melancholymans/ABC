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
	r := strings.Split(sc.Text(), " ")
	a1, _ := strconv.Atoi(r[0])
	a2, _ := strconv.Atoi(r[1])
	a3, _ := strconv.Atoi(r[2])
	ad := make([]int, 3)
	ad[0] = Abs(a1 - a2)
	ad[1] = Abs(a2 - a3)
	ad[2] = Abs(a3 - a1)
	sort.Sort(sort.IntSlice(ad))
	fmt.Fprintln(writer, ad[0]+ad[1])
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
