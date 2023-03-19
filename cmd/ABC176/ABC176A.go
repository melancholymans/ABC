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
	s := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(s[0])
	x, _ := strconv.Atoi(s[1])
	t, _ := strconv.Atoi(s[2])
	fmt.Fprintln(writer, Ceil(float64(n)/float64(x))*t)
}

func Ceil(f float64) int {
	return int(math.Ceil(f))
}
