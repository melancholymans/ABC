package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	m := map[byte]int{
		97: 2, 98: 5, 99: 7, 100: 11, 101: 13, 102: 17, 103: 19, 104: 23, 105: 29, 106: 31, 107: 37, 108: 41, 109: 43, 110: 47, 111: 53, 112: 59, 113: 61, 114: 67, 115: 71, 116: 73, 117: 79, 118: 83, 119: 89, 120: 97, 121: 101, 122: 103,
	}
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 100000)
	sc.Buffer(buf, 100000)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	ad := make([]int, 100000)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		total := 0
		sc.Scan()
		s := []byte(sc.Text())
		for j := 0; j < 10; j++ {
			total += m[s[j]]
		}
		ad[i] = total
	}
	var count int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if ad[i] == ad[j] {
				count += 1
			}
		}
	}
	fmt.Fprintln(writer, count)
}
