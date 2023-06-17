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
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	s := sc.Text()
	sl := make([]string, n*2)
	for i := 0; i < n; i++ {
		sl[i*2] = string(s[i])
		sl[i*2+1] = string(s[i])
	}
	fmt.Fprintln(writer, strings.Join([]string(sl), ""))
}
