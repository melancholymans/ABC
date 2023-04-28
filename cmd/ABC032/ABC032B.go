package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	s := sc.Text()
	sc.Scan()
	k, _ := strconv.Atoi(sc.Text())
	m := map[string]bool{}
	for i := 0; i < len(s)-k+1; i++ {
		if _, ok := m[s[i:k+i]]; ok {
			continue
		}
		m[s[i:k+i]] = true
	}
	fmt.Fprintln(writer, len(m))
}
