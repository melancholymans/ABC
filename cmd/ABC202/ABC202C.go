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
csはindex情報で、cmはそのindexが何個あるかをカウントしている
入力例 1
as 1 2 2
bs 3 1 2
cs 2 3 2
cmは[2:2,3:1]となる
bs[0+1]はcm側にはない。なのでas[1]==bs[1]となることはない
bs[1+1]はcm側に1回出てくる。つまりas[1]==bs[2]となる
bs[2+1]はcm側に2回出てくる。つまりas[2]==as[3]==bs[3]となるので
countは１＋２＋２＝４となる
*/
func main() {
	defer wtr.Flush()
	n := ni()
	as := nis(n)
	bs := nis(n)
	cs := nis(n)
	count := 0
	cm := make(map[int]int)
	for _, v := range cs {
		if _, ok := cm[v]; ok {
			cm[v]++
		} else {
			cm[v] = 1
		}
	}
	bm := make(map[int]int)
	for i, v := range bs {
		t := 0
		if _, ok := cm[i+1]; ok {
			t = cm[i+1]
		}
		if _, ok := bm[v]; ok {
			bm[v] += t
		} else {
			bm[v] = t
		}
	}
	for _, v := range as {
		if _, ok := bm[v]; ok {
			count += bm[v]
		}
	}
	fmt.Fprintln(wtr, count)
}
