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
	n, _ := strconv.Atoi(sc.Text())
	sl := make([][]string, n)
	for i := 0; i < n; i++ {
		sl[i] = make([]string, n)
		sc.Scan()
		r := strings.Split(sc.Text(), "")
		for j := 0; j < n; j++ {
			sl[i][j] = r[j]
		}
	}
	ad := turnLeft(n, sl)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprint(writer, ad[i][j])
		}
		fmt.Fprint(writer, "\n")
	}
}

// 正行列、行数＝列数が前提
func turnRight(n int, sl [][]string) [][]string {
	ad := make([][]string, n)
	for i := 0; i < n; i++ {
		ad[i] = make([]string, n)
		for j := 0; j < n; j++ {
			ad[i][j] = sl[n-1-j][i]
		}
	}
	return ad
}

func turnLeft(n int, sl [][]string) [][]string {
	ad := make([][]string, n)
	for i := 0; i < n; i++ {
		ad[i] = make([]string, n)
		for j := 0; j < n; j++ {
			ad[i][j] = sl[j][n-1-i]
		}
	}
	return ad
}

/*
func Sli(l, m int) [][]string {
	sl := make([][]string, l)
	for i := 0; i < l; i++ {
		sl[i] = make([]string, m)
		for j := 0; j < m; j++ {
			sl[i][j] = ""
		}
	}
	return sl
}
*/
