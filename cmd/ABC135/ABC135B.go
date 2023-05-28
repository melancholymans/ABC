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
	r1 := strings.Split(sc.Text(), " ")
	sl := make([]int, 0)
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(r1[i])
		sl = append(sl, a)
	}
	slc := make([]int, len(sl))
	copy(slc, sl)
	sort.Sort(sort.IntSlice(slc))
	count := 0
	for i := 0; i < n; i++ {
		if sl[i] != slc[i] {
			count += 1
		}
	}
	if count > 2 {
		fmt.Fprintln(writer, "NO")
	} else {
		fmt.Fprintln(writer, "YES")
	}
}
