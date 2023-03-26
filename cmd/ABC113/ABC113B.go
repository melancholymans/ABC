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
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	s := strings.Split(sc.Text(), " ")
	t, _ := strconv.ParseFloat(s[0], 64)
	a, _ := strconv.ParseFloat(s[1], 64)
	sc.Scan()
	s = strings.Split(sc.Text(), " ")
	var result int
	tmin := math.MaxFloat64
	for i := 0; i < n; i++ {
		h, _ := strconv.ParseFloat(s[i], 64)
		tn := math.Abs(a - (t - h*0.006))
		if tmin > tn {
			tmin = tn
			result = i
		}
	}
	fmt.Fprintln(writer, result+1)
}
