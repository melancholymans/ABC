package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func substr(s string, i int, l int) string {
	if i+l > len(s) {
		return ""
	}
	return s[i : i+l]
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0)
	sc.Buffer(buf, 131072)
	sc.Scan()
	s := sc.Text()
	rs := reverse(s)
	divide := []string{"maerd", "remaerd", "esare", "resare"}
	var can = true
	var can2 bool
	i := 0
	limit := len(s)
	for i < limit {
		can2 = false
		for j := 0; j < 4; j++ {
			d := divide[j]
			if substr(rs, i, len(d)) == d {
				can2 = true
				i += len(d)
			}
			if i >= limit {
				break
			}
		}
		if !can2 {
			can = false
			break
		}
	}
	if can {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
