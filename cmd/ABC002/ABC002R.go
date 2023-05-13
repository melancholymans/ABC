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

/*
	fmt.Println("i=", i, "j=", j, "sl[j][w-1-i]=", sl[j][w-1-i], "w-1-i=", w-1-i)
*/
