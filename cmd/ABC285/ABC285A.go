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
func main() {
	defer wtr.Flush()
	a, b := ni2()
	var sl [16][16]int
	sl[1][2] = 1
	sl[1][3] = 1
	sl[2][1] = 1
	sl[3][1] = 1
	sl[2][4] = 1
	sl[2][5] = 1
	sl[4][2] = 1
	sl[5][2] = 1
	sl[3][6] = 1
	sl[3][7] = 1
	sl[6][3] = 1
	sl[7][3] = 1
	sl[4][8] = 1
	sl[4][9] = 1
	sl[8][4] = 1
	sl[9][4] = 1
	sl[5][10] = 1
	sl[5][11] = 1
	sl[10][5] = 1
	sl[11][5] = 1
	sl[6][12] = 1
	sl[6][13] = 1
	sl[12][6] = 1
	sl[13][6] = 1
	sl[7][14] = 1
	sl[7][15] = 1
	sl[14][7] = 1
	sl[15][7] = 1
	if sl[a][b] == 1 {
		fmt.Fprintln(wtr, "Yes")
	} else {
		fmt.Fprintln(wtr, "No")
	}
}
