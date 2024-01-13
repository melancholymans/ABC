package main

import (
	"bufio"
	"container/heap"
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

type graph struct {
	size         int
	edges        [][]edge
	starts       []state
	comps        []compFunc
	defaultScore int
	level        []int
	iter         []int
}

type pq struct {
	arr   []interface{}
	comps []compFunc
}

func main() {
	defer wtr.Flush()
	total := 0
	v, e := ni2()
	edges := make([][]edge, v)

	for i := 0; i < e; i++ {
		s := ni() - 1
		t := ni() - 1
		edges[s] = append(edges[s], edge{to: t, cost: 1})
	}
	for i := 0; i < v; i++ {
		graph := newGraph(v, edges)
		graph.setStart(state{node: i})
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
		dist := graph.dijkstra()
		for _, v := range dist {
			if v < math.MaxInt64 {
				total++
			}
		}
	}
	fmt.Fprintln(wtr, total)
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

func (g *graph) setStart(s state) {
	g.starts = append(g.starts, s)
}

func (g *graph) dijkstra() []int {
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

func (pq pq) Len() int {
	return len(pq.arr)
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

func (pq *pq) IsEmpty() bool {
	return pq.Len() == 0
}

func (pq *pq) Pop() interface{} {
	n := pq.Len()
	item := pq.arr[n-1]
	pq.arr = pq.arr[:n-1]
	return item
}

func (pq *pq) Push(x interface{}) {
	pq.arr = append(pq.arr, x)
}

func (pq pq) Swap(i, j int) {
	pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
}
