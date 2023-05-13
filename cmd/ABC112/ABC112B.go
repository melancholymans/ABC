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
	tld, _ := strconv.Atoi(r1[1])
	minm := math.MaxInt64
	for i := 0; i < n; i++ {
		sc.Scan()
		r2 := strings.Split(sc.Text(), " ")
		c, _ := strconv.Atoi(r2[0])
		t, _ := strconv.Atoi(r2[1])
		if t > tld {
			continue
		}
		if minm > c {
			minm = c
		}
	}
	if minm == math.MaxInt64 {
		fmt.Fprintln(writer, "TLE")
	} else {
		fmt.Fprintln(writer, minm)
	}
}
