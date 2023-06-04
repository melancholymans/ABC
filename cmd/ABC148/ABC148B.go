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
	r := strings.Split(sc.Text(), " ")
	s := r[0]
	t := r[1]
	sl := make([]string, 0)
	for i := 0; i < n; i++ {
		sl = append(sl, string(s[i]))
		sl = append(sl, string(t[i]))
	}
	fmt.Fprintln(writer, strings.Join(sl, ""))
}
