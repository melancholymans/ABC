package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type IntPair [2]float64

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	s := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(s[0])
	d, _ := strconv.ParseFloat(s[1], 64)
	ad := make([]IntPair, n)
	count := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		la := strings.Split(sc.Text(), " ")
		ad[i][0], _ = strconv.ParseFloat(la[0], 64)
		ad[i][1], _ = strconv.ParseFloat(la[1], 64)
		dis := math.Sqrt(math.Pow(ad[i][0], 2) + math.Pow(ad[i][1], 2))
		if dis <= d {
			count += 1
		}
	}
	fmt.Fprintln(writer, count)
}
