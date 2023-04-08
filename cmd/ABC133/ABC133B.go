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
	r1 := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(r1[0])
	d, _ := strconv.Atoi(r1[1])
	sl := make([][]float64, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		r2 := strings.Split(sc.Text(), " ")
		sl[i] = make([]float64, d)
		for j := 0; j < d; j++ {
			a, _ := strconv.Atoi(r2[j])
			sl[i][j] = float64(a)
		}
	}
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			count += distance(sl, i, j, d)
		}
	}
	fmt.Fprintln(writer, count)
}

func distance(sl [][]float64, i, j, d int) int {
	ds := 0.0
	count := 0
	for k := 0; k < d; k++ {
		ds += math.Pow(sl[i][k]-sl[j][k], 2)
	}
	sq := math.Sqrt(ds)
	if math.Floor(sq) == sq {
		count += 1
	}
	return count
}
