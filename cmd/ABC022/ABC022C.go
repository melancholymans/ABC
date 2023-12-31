package main

import (
	"bufio"
	"container/heap"
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

func main() {
	defer wtr.Flush()
	mmin := math.MaxInt64
	n, m := ni2()
	edges := make([][]edge, n)
	is := make([]int, n)
	for i := 0; i < m; i++ {
		u := ni() - 1
		v := ni() - 1
		l := ni()
		edges[u] = append(edges[u], edge{to: v, cost: l})
		edges[v] = append(edges[v], edge{to: u, cost: l})
		if v == 0 {
			is[u] = len(edges[u]) - 1
		}
		if u == 0 {
			is[v] = len(edges[v]) - 1
		}
	}
	for _, nv := range edges[0] {
		tc := edges[nv.to][is[nv.to]].cost
		edges[nv.to][is[nv.to]].cost = int(1e10)
		graph := newGraph(n, edges)
		dist := graph.dijkstra(nv.to)
		mmin = min(mmin, nv.cost+dist[0])
		edges[nv.to][is[nv.to]].cost = tc
	}
	if mmin >= int(1e10) {
		fmt.Fprintln(wtr, -1)
		return
	}
	fmt.Fprintln(wtr, mmin)
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

func (g *graph) dijkstra(start int) []int {
	g.starts = []state{{node: start}}
	score := make([]int, g.size)
	for i := 0; i < g.size; i++ {
		score[i] = g.defaultScore
	}
	que := newpq(g.comps)
	for _, start := range g.starts {
		score[start.node] = start.score
		heap.Push(que, start)
	}
	for !que.IsEmpty() {
		st := heap.Pop(que).(state)
		if st.score > score[st.node] {
			continue
		}
		for _, edge := range g.edges[st.node] {
			newScore := st.score + edge.cost
			if score[edge.to] > newScore {
				score[edge.to] = newScore
				heap.Push(que, state{score: newScore, node: edge.to})
			}
		}
	}
	return score
}

func newpq(comps []compFunc) *pq {
	return &pq{
		comps: comps,
	}
}

func (pq *pq) Push(x interface{}) {
	pq.arr = append(pq.arr, x)
}

func (pq *pq) Pop() interface{} {
	n := pq.Len()
	item := pq.arr[n-1]
	pq.arr = pq.arr[:n-1]
	return item
}
func (pq *pq) IsEmpty() bool {
	return pq.Len() == 0
}

func (pq pq) Len() int {
	return len(pq.arr)
}

func (pq pq) Swap(i, j int) {
	pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
}

func (pq pq) Less(i, j int) bool {
	for _, comp := range pq.comps {
		result := comp(pq.arr[i], pq.arr[j])
		switch result {
		case -1:
			return true
		case 1:
			return false
		case 0:
			continue
		}
	}
	return true
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
