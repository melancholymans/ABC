package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	r1 := strings.Split(sc.Text(), " ")
	h, _ := strconv.Atoi(r1[0])
	w, _ := strconv.Atoi(r1[1])
	sl := make([][]string, h)
	for i := 0; i < h; {
		sc.Scan()
		r2 := strings.Split(sc.Text(), "")
		if calc(r2) {
			sl[i] = append(sl[i], r2...)
		}
		i += 1
	}
	fmt.Fprintln(writer, sl, "\n", w)
}

func calc(r []string) bool {
	for i := 0; i < len(r); i++ {
		if r[i] == "#" {
			return true
		}
	}
	return false
}

/*
func Sli(l, m int, def int) [][]int {
	sl := make([][]int, l)
	for i := 0; i < l; i++ {
		sl[i] = make([]int, m)
		for j := 0; j < m; j++ {
			sl[i][j] = def
		}
	}
	return sl
}
*/
/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	//sc.Scan()
	//n, _ := strconv.Atoi(sc.Text())
	h := 2
	w := 4
	sl := make([][]string, h)
	for i := 0; i < h; i++ {
		sl[i] = make([]string, w)
		sc.Scan()
		r := strings.Split(sc.Text(), " ")
		for j := 0; j < w; j++ {
			sl[i][j] = r[j]
		}
	}
	ad := turnRight(h, w, sl)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			fmt.Fprint(writer, ad[i][j], " ")
		}
		fmt.Fprint(writer, "\n")
	}
}

func turnRight(h, w int, sl [][]string) [][]string {
	ad := make([][]string, w)
	for i := 0; i < w; i++ {
		ad[i] = make([]string, h)
		for j := 0; j < h; j++ {
			ad[i][j] = sl[h-1-j][i]
		}
	}
	return ad
}

func turnLeft(h, w int, sl [][]string) [][]string {
	ad := make([][]string, w)
	for i := 0; i < w; i++ {
		ad[i] = make([]string, h)
		for j := 0; j < h; j++ {
			ad[i][j] = sl[j][w-1-i]
		}
	}
	return ad
}
*/
