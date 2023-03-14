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
	m := make(map[string]int, 100000)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		sc.Scan()
		s := strings.Split(sc.Text(), "")
		sort.Strings(s)
		curs := strings.Join(s, "")
		m[curs] += 1
	}
	result := 0
	for _, v := range m {
		result += v * (v - 1) / 2
	}
	fmt.Fprintln(writer, result)
}
