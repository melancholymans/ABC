package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := ioutil.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}
}
func ni() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
func nis(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ni()
	}
	return a
}
func ni2() (int, int) {
	return ni(), ni()
}
func ni3() (int, int, int) {
	return ni(), ni(), ni()
}
func ni4() (int, int, int, int) {
	return ni(), ni(), ni(), ni()
}

func ns() string {
	sc.Scan()
	return sc.Text()
}

func nf() float64 {
	sc.Scan()
	f, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return f
}

/*
###.?????
###.?????
###.?????
....?????
?????????
?????....
?????.###
?????.###
?????.###
*/
/*
? byte(63)
*/
func main() {
	defer wtr.Flush()
	s := `###.?????###.?????###.?????....???????????????????....?????.###?????.###?????.###`
	base := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		base[i] = make([]byte, 9)
		for j := 0; j < 9; j++ {
			base[i][j] = s[j*9+i]
		}
	}
	n, m := ni2()
	sl := make([][]byte, n)
	for i := 0; i < n; i++ {
		sl[i] = []byte(ns())
	}
	for i := 0; i <= n-9; i++ {
		for j := 0; j <= m-9; j++ {
			flag := true
			for k := 0; k < 9; k++ {
				for l := 0; l < 9; l++ {
					if base[k][l] == byte(63) || base[k][l] == sl[i+k][j+l] {
					} else {
						flag = false
					}
				}
			}
			if flag {
				fmt.Fprintf(wtr, "%d %d\n", i+1, j+1)
			}
		}
	}
}
