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
	a, _ := strconv.Atoi(s[0])
	b, _ := strconv.Atoi(s[1])
	calc(writer, a, b)
}

func calc(writer *bufio.Writer, a, b int) {
	for i := 1; i <= 1250; i++ {
		if Floor(float64(i)*float64(0.08)) == a && Floor(float64(i)*float64(0.1)) == b {
			fmt.Fprintln(writer, i)
			return
		}
	}
	fmt.Fprintln(writer, -1)
}

func Floor(f float64) int {
	return int(math.Floor(f))
}
