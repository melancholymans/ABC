package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	ad := make([]int, 0)
	for i := 0; i < n; i++ {
		sc.Scan()
		r, _ := strconv.Atoi(sc.Text())
		ad = append(ad, r)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ad)))
	sig := 1.0
	total := 0.0
	for i := 0; i < n; i++ {
		total += sig * float64(ad[i]*ad[i])
		sig = sig * -1.0
	}
	fmt.Fprintln(writer, total*math.Pi)
}
