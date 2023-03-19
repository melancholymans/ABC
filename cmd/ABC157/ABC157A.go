package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	s := sc.Text()
	n, _ := strconv.Atoi(s)
	fmt.Fprintln(writer, Ceil(float64(n)/2.0))
}

func Ceil(f float64) int {
	return int(math.Ceil(f))
}
