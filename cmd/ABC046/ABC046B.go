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
	k, _ := strconv.ParseFloat(r[1], 64)
	fmt.Fprintln(writer, int(k*math.Pow(k-1, n-1)))
}
