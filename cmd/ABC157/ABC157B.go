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
	var ad [3][3]int
	for i := 0; i < 3; i++ {
		sc.Scan()
		r := strings.Split(sc.Text(), " ")
		for j := 0; j < 3; j++ {
			a, _ := strconv.Atoi(r[j])
			ad[i][j] = a
		}
	}
	var c [3][3]int
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		sc.Scan()
		b, _ := strconv.Atoi(sc.Text())
		if x, y, ok := Contains(ad, b); ok {
			c[x][y] = 1
		}
	}
	for i := 0; i < 3; i++ {
		count := 0
		for j := 0; j < 3; j++ {
			count += c[i][j]
		}
		if count == 3 {
			fmt.Fprintln(writer, "Yes")
			return
		}
	}
	for j := 0; j < 3; j++ {
		count := 0
		for i := 0; i < 3; i++ {
			count += c[i][j]
		}
		if count == 3 {
			fmt.Fprintln(writer, "Yes")
			return
		}
	}
	if c[0][0]+c[1][1]+c[2][2] == 3 {
		fmt.Fprintln(writer, "Yes")
		return
	}
	if c[0][2]+c[1][1]+c[0][2] == 3 {
		fmt.Fprintln(writer, "Yes")
		return
	}
	fmt.Fprintln(writer, "No")
}

func Contains(ad [3][3]int, f int) (i, j int, b bool) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if ad[i][j] == f {
				return i, j, true
			}
		}
	}
	return 0, 0, false
}
