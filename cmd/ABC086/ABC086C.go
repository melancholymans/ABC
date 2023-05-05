package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0)
	sc.Buffer(buf, 131072)
	sc.Split(bufio.ScanWords)

	var a [3][110000]int
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	for c := 1; c < n+1; c++ {
		for r := 0; r < 3; r++ {
			sc.Scan()
			a[r][c], _ = strconv.Atoi(sc.Text())
		}
	}
	flag := true
	for i := 0; i < n; i++ {
		dt := a[0][i+1] - a[0][i]
		dst := abs(a[1][i+1]-a[1][i]) + abs(a[2][i+1]-a[2][i])
		if dst > dt {
			flag = false
		}
		if dst%2 != dt%2 {
			flag = false
		}
	}
	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
