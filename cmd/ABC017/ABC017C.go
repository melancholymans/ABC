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

type bit struct {
	n int
	b []int
}

func newBit(n int) *bit {
	return &bit{
		n: n + 1,
		b: make([]int, n+1),
	}
}

func (b *bit) add(i, x int) {
	for i++; i < b.n && i > 0; i += i & -i {
		b.b[i] += x
	}
}

func (b *bit) rangesum(l, r int) int {
	return b.sum(r-1) - b.sum(l-1)
}

func (b *bit) sum(i int) int {
	ret := 0
	for i++; i > 0; i -= i & -i {
		ret += b.b[i]
	}
	return ret
}

type taple [3]int

func main() {
	defer wtr.Flush()
	n, m := ni2()
	sl := make([]taple, n)
	for i := 0; i < n; i++ {
		sl[i][0], sl[i][1], sl[i][2] = ni(), ni(), ni()
	}
	fmt.Fprintln(wtr, n, m)
	fmt.Fprintln(wtr, sl)
}

func main() {

	defer flush()

	o := 0
	n, m := ni2()
	lrs := make([][][2]int, m)
	sum := 0
	for i := 0; i < n; i++ {
		l, r, s := ni3()
		l--
		r--
		lrs[l] = append(lrs[l], [2]int{r, s})
		sum += s
	}
	bit := newBit(m)
	for i := 0; i < m; i++ {
		for _, v := range lrs[i] {
			bit.add(v[0], v[1])
		}
		maxs(&o, sum-bit.rangesum(i, m))
	}

	out(o)
}
