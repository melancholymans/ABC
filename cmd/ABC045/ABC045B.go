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
	card := make([][]string, 3)
	sc.Scan()
	card[0] = strings.Split(sc.Text(), "")
	sc.Scan()
	card[1] = strings.Split(sc.Text(), "")
	sc.Scan()
	card[2] = strings.Split(sc.Text(), "")
	m := map[string]int{"a": 0, "b": 1, "c": 2}
	m2 := map[int]string{0: "A", 1: "B", 2: "C"}
	start := 0
	for {
		if len(card[start]) == 0 {
			fmt.Fprintln(writer, m2[start])
			return
		}
		s := card[start][0]
		card[start] = append(card[start][1:])
		start = m[s]
	}
}
