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
	m := map[int]int{
		1: 0,
		2: 1,
		3: 0,
		4: 1,
		5: 2,
		6: 3,
		7: 0,
		8: 1,
		9: 0,
	}
	count := 0
	sc.Scan()
	r1 := strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(r1[i])
		count += m[a]
	}
	fmt.Fprintln(writer, count)
}
