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
	n, _ := strconv.ParseFloat(r[0], 64)
	m, _ := strconv.ParseFloat(r[1], 64)
	k := math.Abs((math.Mod(n, 12.0)+m/60.0)*30 - m*6)
	fmt.Fprintln(writer, math.Min(k, 360-k))
}
