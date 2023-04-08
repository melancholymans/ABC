package main

import (
	"bufio"
	"fmt"
	"math"
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
	n, _ := strconv.Atoi(r1[0])
	d, _ := strconv.Atoi(r1[1])
	sl := make([][]float64, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		r2 := strings.Split(sc.Text(), " ")
		sl[i] = make([]float64, d)
		for j := 0; j < d; j++ {
			a, _ := strconv.Atoi(r2[j])
			sl[i][j] = float64(a)
		}
	}
	fmt.Println("-----------")
	var ds float64
	for j := 0; j < d; j++ {
		ds += math.Pow(sl[0][j]-sl[1][j], 2)
	}
	sq := math.Sqrt(ds)
	if math.Floor(sq) == sq {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	ds = 0.0
	for j := 0; j < d; j++ {
		ds += math.Pow(sl[1][j]-sl[2][j], 2)
	}
	sq = math.Sqrt(ds)
	if math.Floor(sq) == sq {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	ds = 0.0
	for j := 0; j < d; j++ {
		ds += math.Pow(sl[2][j]-sl[0][j], 2)
	}
	sq = math.Sqrt(ds)
	if math.Floor(sq) == sq {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
