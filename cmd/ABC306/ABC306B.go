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
	r := strings.Split(sc.Text(), " ")
	sl := make([]int, 64)
	for i := 0; i < 64; i++ {
		a, _ := strconv.Atoi(r[i])
		sl[i] = a
	}
	total := 0.0
	fmt.Fprintln(writer, sl)

	for i := 0; i < 64; i++ {
		total += math.Pow(2.0, float64(i)) * float64(sl[i])
	}
	fmt.Fprintln(writer, int64(total))
}
