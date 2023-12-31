package main

import (
	"bufio"
	"container/list"
	"fmt"
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
		b, e := os.ReadFile("./input")
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

type edge struct {
	from int
	to   int
	cost int
	rev  int
}

type state struct {
	score int
	node  int
}

type compFunc func(p, q interface{}) int

type pq struct {
	arr   []interface{}
	comps []compFunc
}

type graph struct {
	size         int
	edges        [][]edge
	starts       []state
	comps        []compFunc
	defaultScore int
	level        []int
	iter         []int
}

type point struct {
	x int
	y int
}

func main() {
	defer wtr.Flush()
	n := ni()
	a := ni() - 1
	b := ni() - 1
	m := ni()
	edges := make([][]edge, n)
	pos := make([]int, n)
	dist := make([]int, n)
	for i := 0; i < m; i++ {
		s := ni() - 1
		t := ni() - 1
		edges[s] = append(edges[s], edge{to: t})
		edges[t] = append(edges[t], edge{to: s})
	}
	//container/listは双方向のリンクリスト
	q := list.New()
	q.PushBack(a)
	pos[a] = 1
	dist[a] = 1
	e := q.Front()
	for e != nil {
		v := e.Value.(int)
		for _, nv := range edges[v] {
			if dist[nv.to] != 0 && (dist[nv.to] < dist[v]+1) {
				continue
			}
			if dist[nv.to] == 0 {
				dist[nv.to] = dist[v] + 1
				q.PushBack(nv.to)
			}
			pos[nv.to] = madd(pos[nv.to], pos[v])
		}
		e = e.Next()
	}
	fmt.Fprintln(wtr, pos[b])
}

const mod = 1000000007

func madd(a, b int) int {
	a += b
	if a < 0 {
		a += mod
	} else if a >= mod {
		a -= mod
	}
	return a
}
