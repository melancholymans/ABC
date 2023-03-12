package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntPair [2]int

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	in := strings.Split(sc.Text(), " ")
	h, _ := strconv.Atoi(in[0])
	w, _ := strconv.Atoi(in[1])
	var cs string
	for i := 0; i < h; i++ {
		sc.Scan()
		s := sc.Text()
		cs = cs + s
	}
	cb := []byte(cs)
	v := [8]IntPair{
		{-1, 0},  //up
		{1, 0},   //down
		{0, -1},  //left
		{0, 1},   //right
		{-1, -1}, //upleft
		{-1, 1},  //upright
		{1, -1},  //downleft
		{1, 1},   //downright
	}
	ad := Sli(h, w, 0) //35=#,46=.
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ad[i][j] = cb[i*w+j]
		}
	}

	rt := Sli(h, w, 0)
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			a := ad[row][col]
			if a == 35 {
				rt[row][col] += 100
			} else {
				for i := 0; i < 8; i++ {
					ncol := col + v[i][1]
					if ncol < 0 || ncol > w-1 {
						continue
					}
					nrow := row + v[i][0]
					if nrow < 0 || nrow > h-1 {
						continue
					}
					b := ad[nrow][ncol]
					if b == 35 {
						rt[row][col] += 1
					}
				}
			}
		}
	}
	var c string
	for i := 0; i < h; i++ {
		cs = ""
		for j := 0; j < w; j++ {
			if rt[i][j] == 100 {
				c = "#"
			} else {
				c = strconv.Itoa(int(rt[i][j]))
			}
			cs = cs + c
		}
		fmt.Fprintln(writer, cs)
	}
}

func Sli(l, m int, def byte) [][]byte {
	sl := make([][]byte, l)
	for i := 0; i < l; i++ {
		sl[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			sl[i][j] = def
		}
	}
	return sl
}
