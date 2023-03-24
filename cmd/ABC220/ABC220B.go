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
	s1 := sc.Text()
	k, _ := strconv.Atoi(s1)
	sc.Scan()
	s2 := strings.Split(sc.Text(), " ")
	a, _ := strconv.Atoi(s2[0])
	b, _ := strconv.Atoi(s2[1])
	//aを10進数にする
	var ar1 int
	lim := int(math.Log10(float64(a))) + 1
	for i := 0; i < lim; i++ {
		r := a % 10
		a = a / 10
		ar1 = ar1 + r*int(math.Pow(float64(k), float64(i)))
	}
	//bを10進数にする
	var ar2 int
	lim = int(math.Log10(float64(b))) + 1
	for i := 0; i < lim; i++ {
		r := b % 10
		b = b / 10
		ar2 = ar2 + r*int(math.Pow(float64(k), float64(i)))
	}
	fmt.Fprintln(writer, ar1*ar2)
}
