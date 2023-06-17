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
	k, _ := strconv.Atoi(r1[1])
	mp := map[int]int{}
	for i := 0; i < k; i++ {
		sc.Scan()
		d, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		r2 := strings.Split(sc.Text(), " ")
		for j := 0; j < d; j++ {
			a, _ := strconv.Atoi(r2[j])
			mp[a] = 1
		}
	}
	fmt.Fprintln(writer, n-len(mp))
}
