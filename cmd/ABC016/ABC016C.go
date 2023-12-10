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

type compFunc func(p, q interface{}) int

type state struct {
	score int
	node  int
}

type edge struct {
	from int
	to   int
	cost int
	rev  int
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

func main() {
	defer wtr.Flush()
	n, m := ni2()
	edges := make([][]edge, n)
	for i := 0; i < m; i++ {
		s := ni() - 1
		t := ni() - 1
		c := 1
		edges[s] = append(edges[s], edge{to: t, cost: c})
		edges[t] = append(edges[t], edge{to: s, cost: c})
	}
	g := newGraph(n, edges)
	dist, _ := g.floydWarshall()
	for i := 0; i < n; i++ {
		t := 0
		for j := 0; j < n; j++ {
			if dist[i][j] == 2 {
				t += 1
			}
		}
		fmt.Fprintln(wtr, t)
	}
}

func newGraph(size int, edges [][]edge) *graph {
	graph := &graph{
		size:  size,
		edges: edges,
	}
	graph.defaultScore = math.MaxInt64
	graph.comps = []compFunc{
		func(p, q interface{}) int {
			if p.(state).score < q.(state).score {
				return -1
			} else if p.(state).score == q.(state).score {
				return 0
			}
			return 1
		},
	}
	return graph
}

// ワーシャル–フロイド法のアルゴリズムで最短経路問題を解く
func (g *graph) floydWarshall() ([][]int, bool) {
	score := make([][]int, g.size)
	for i := 0; i < g.size; i++ {
		score[i] = make([]int, g.size)
		for j := 0; j < g.size; j++ {
			if i == j {
				score[i][j] = 0
			} else {
				score[i][j] = g.defaultScore
			}
		}
		for _, edge := range g.edges[i] {
			score[i][edge.to] = edge.cost
		}
	}
	for k := 0; k < g.size; k++ {
		for i := 0; i < g.size; i++ {
			for j := 0; j < g.size; j++ {
				if score[i][k] == g.defaultScore || score[k][j] == g.defaultScore {
					continue
				}
				score[i][j] = min(score[i][j], score[i][k]+score[k][j])
			}
		}
	}
	for k := 0; k < g.size; k++ {
		if score[k][k] < 0 {
			return nil, true
		}
	}
	return score, false
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
