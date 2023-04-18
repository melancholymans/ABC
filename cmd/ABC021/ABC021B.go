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
	m := map[string]int{}
	sc.Scan()
	//n, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	r1 := strings.Split(sc.Text(), " ")
	a := r1[0]
	b := r1[1]
	m[a] += 1
	m[b] += 1
	sc.Scan()
	k, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")

	for i := 0; i < k; i++ {
		p := r2[i]
		m[p] += 1
	}
	for _, v := range m {
		if v > 1 {
			fmt.Fprintln(writer, "NO")
			return
		}
	}
	fmt.Fprintln(writer, "YES")
}
