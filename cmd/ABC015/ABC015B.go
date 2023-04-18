package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	r1 := sc.Text()
	n, _ := strconv.Atoi(r1)
	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	ad := make([]float64, 0)
	for i := 0; i < n; i++ {
		a, _ := strconv.ParseFloat(r2[i], 64)
		ad = append(ad, a)
	}
	count := 0.0
	var total float64
	for i := 0; i < n; i++ {
		if ad[i] > 0 {
			count += 1.0
			total += ad[i]
		}
	}
	fmt.Fprintln(writer, math.Ceil(total/count))
}
