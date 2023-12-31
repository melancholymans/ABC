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

type point struct {
	x int
	y int
}

func main() {
	defer wtr.Flush()
	h, w, t := ni3()
	sl := make([][]int, h)
	chg := map[byte]int{46: 0, 35: 1, 83: 2, 71: 3} //46=".",35="#",83=S,71="G"
	for i := 0; i < h; i++ {
		bl := []byte(ns())
		sl[i] = make([]int, w)
		for j := 0; j < w; j++ {
			sl[i][j] = chg[bl[j]]
		}
	}
	var st, gl int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if sl[i][j] == 2 {
				sl[i][j] = 0
				st = i*w + j
			}
			if sl[i][j] == 3 {
				sl[i][j] = 0
				gl = i*w + j
			}
		}
	}
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	calc := func(x int) bool {
		edges := make([][]edge, h*w)
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				for k := 0; k < 4; k++ {
					p := point{i, j}
					np := pointAdd(p, point{dx[k], dy[k]}) //np->next point
					if !np.isValid(h, w) {
						continue
					}
					s := p.x*w + p.y
					t := np.x*w + np.y
					c := 1
					if sl[np.x][np.y] == 1 {
						c = x //npの値が#==1なら渡されたxをコストとして記録する、そうでなけれな1のままとする
					}
					edges[s] = append(edges[s], edge{to: t, cost: c}) //ノードを追加していく
				}
			}
		}
		graph := newGraph(h*w, edges)
		dist := graph.dijkstra(st) //ダイクストラ法
		return dist[gl] <= t
	}
	fmt.Fprintln(wtr, bs(1, int(1e9), calc))
}

func bs(ok, ng int, f func(int) bool) int {
	if !f(ok) {
		return -1
	}
	if f(ng) {
		return ng
	}
	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2

		if f(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func pointAdd(a, b point) point {
	return point{x: a.x + b.x, y: a.y + b.y}
}

func (p point) isValid(h, w int) bool {
	return (0 <= p.x && p.x < h) && (0 <= p.y && p.y < w)
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

/*
"container/heap"をベースにしたヒープデータ構造
２分木ソートもできる。このプログラムのなかではダイクストラ法で
グラフを実装するのに使用されている
*/
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
