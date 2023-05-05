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
	n, _ := strconv.Atoi(sc.Text())
	for i := n; i > 0; i-- {
		x, _ := Divisor(i)
		rst, ok := calc(x)
		if ok {
			fmt.Fprintln(writer, rst)
			return
		}
	}
}

func calc(x []int) (int, bool) {
	l := len(x)
	if l%2 != 0 {
		return 0, false
	}
	s := x[0]
	for i := 1; i < l; i++ {
		if s != x[i] {
			return 0, false
		}
	}
	for i := 0; i < l/2; i++ {
		s = s * s
	}
	return s, true
}

func Divisor(n int) ([]int, map[int]int) {
	sqrtn := int(math.Sqrt(float64(n)))
	c := 2
	divisor := []int{}
	divisorm := make(map[int]int)
	for {
		if n%2 != 0 {
			break
		}
		divisor = append(divisor, 2)
		divisorm[2]++
		n /= 2
	}
	c = 3
	for {
		if n%c == 0 {
			divisor = append(divisor, c)
			divisorm[c]++
			n /= c
		} else {
			c += 2
			if c > sqrtn {
				break
			}
		}
	}
	if n != 1 {
		divisor = append(divisor, n)
		divisorm[n]++
	}
	return divisor, divisorm
}
