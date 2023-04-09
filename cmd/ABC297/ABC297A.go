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
	r1 := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(r1[0])
	d, _ := strconv.Atoi(r1[1])
	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	sl := make([]int, 0)
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(r2[i])
		sl = append(sl, a)
	}
	for i := 0; i < len(sl)-1; i++ {
		if sl[i+1]-sl[i] <= d {
			fmt.Fprintln(writer, sl[i+1])
			return
		}
	}
	fmt.Fprintln(writer, -1)
}
