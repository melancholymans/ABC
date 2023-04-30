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
	n, _ := strconv.Atoi(r[0])
	strconv.Atoi(r[1])
	sl := make([]string, 0)
	for i := 0; i < n; i++ {
		sc.Scan()
		s := sc.Text()
		sl = append(sl, s)
	}
	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})
	fmt.Fprintln(writer, strings.Join(sl, ""))
}
