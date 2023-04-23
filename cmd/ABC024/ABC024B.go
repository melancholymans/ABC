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
	r := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(r[0])
	t, _ := strconv.Atoi(r[1])
	ad := make([]int, 0)
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		ad = append(ad, a)
	}
	total := 0
	for i := 0; i < n-1; i++ {
		if ad[i+1]-ad[i] > t {
			total += t
		} else {
			total += ad[i+1] - ad[i]
		}
	}
	fmt.Fprintln(writer, total+t)
}
