package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vc struct {
	s string
	d int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	r1 := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(r1[0])
	a, _ := strconv.Atoi(r1[1])
	b, _ := strconv.Atoi(r1[2])
	ad := make([]vc, 0)
	for i := 0; i < n; i++ {
		var v vc
		sc.Scan()
		r2 := strings.Split(sc.Text(), " ")

		v.s = r2[0]
		v.d, _ = strconv.Atoi(r2[1])
		ad = append(ad, v)
	}
	pos := 0
	for i := 0; i < n; i++ {
		if ad[i].s == "East" {
			d := ad[i].d
			if d > b {
				pos += b
			} else if d < a {
				pos += a
			} else {
				pos += d
			}
		} else {
			d := ad[i].d
			if d > b {
				pos -= b
			} else if d < a {
				pos -= a
			} else {
				pos -= d
			}
		}
	}
	if pos < 0 {
		fmt.Fprintln(writer, "West", -pos)
	} else if pos > 0 {
		fmt.Fprintln(writer, "East", pos)
	} else {
		fmt.Fprintln(writer, pos)
	}
}
