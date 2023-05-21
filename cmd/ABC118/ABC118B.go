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
	//m, _ := strconv.Atoi(r1[1])
	mp := map[int]int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		r2 := strings.Split(sc.Text(), " ")
		k, _ := strconv.Atoi(r2[0])
		for j := 0; j < k; j++ {
			a, _ := strconv.Atoi(r2[j+1])
			mp[a] += 1
		}
	}
	count := 0
	for _, v := range mp {
		if v == n {
			count += 1
		}
	}
	fmt.Fprintln(writer, count)
}
