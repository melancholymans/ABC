package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
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

func nis(n int) []int64 {
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		a[i] = int64(ni())
	}
	return a
}

func ni2() (int, int) {
	return ni(), ni()
}

func main() {
	defer wtr.Flush()
	n := ni()
	ns := nis(n)
	rmult := big.NewInt(1)
	e08 := big.NewInt(1000000000000000000)
	for i := 0; i < n; i++ {
		rmult.Mul(rmult, big.NewInt(int64(ns[i])))
	}
	if e08.Cmp(rmult) == -1 {
		fmt.Fprintln(wtr, -1)
	} else {
		fmt.Fprintln(wtr, rmult)
	}
}
