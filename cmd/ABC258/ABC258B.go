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
	dfx:=[8]int{1,0,-1,0,1,1,-1,-1}
	dfy:=[8]int{0,1,0,-1,1,-1,1,-1}
	rst:=""
	n := ni()
	sl := make([][]int, n)
	for i := 0; i < n; i++ {
		sl[i] = make([]int, n)
		for j, s := range strings.Split(ns(), "") {
			num, _ := strconv.Atoi(s)
			sl[i][j] = num
		}
	}
	for lp := 0; lp < n; lp++ {
		for v:=0;v<8;v++{
			x:=lp
			y:=0
			if v==0 || v==2{
				x,y=y,x
			}
			var string s
			for i:=0;i<n;i++{
				s+=strconv.Itoa(sl[x][y])
				x+=dfx[v]
				y+=dfy[v]
				x=(x+n)%n
				y=(y+n)%n
			}
			string t:=s+s
		}
	}
	fmt.Fprintln(wtr, rst)
}

func area(n int, sl [][]int) (int, int, string) {
	mmax := math.MinInt64
	idh, idw := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mmax < sl[i][j] {
				mmax = sl[i][j]
				idh = i
				idw = j
			}
		}
	}
	return idh, idw, strconv.Itoa(mmax)
}

/*
func oneNumber(){

}
*/

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
