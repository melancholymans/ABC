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
func main() {
	defer wtr.Flush()
	n, m, _ := ni3() //t skip
	sl := make([][2]int, m+1)
	for i := 1; i < m+1; i++ {
		a, b := ni2()
		sl[i][0] = a
		sl[i][1] = b
	}
	sl[0][0] = 0
	sl[0][1] = sl[1][0]
	//z := n - (sl[0][0] - 0)
	///if z < 0 {
	//	fmt.Fprintln(wtr, "No")
	//	return
	//}
	z := n
	for i := 0; i < m; i++ {
		//coffeまでの放電処理
		z -= sl[i][1] - sl[i][0]
		//coffeでの充電処理
		z += sl[i+1][1] - sl[i+1][0]
		//ここで放電判定
		if z < 0 {
			fmt.Fprintln(wtr, "No")
			return
		}
		fmt.Fprintln(wtr, "i", i, "z=", z)
	}

}
