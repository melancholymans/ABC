package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	s := []byte(sc.Text())
	sc.Scan()
	t := []byte(sc.Text())
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	sort.Slice(t, func(i, j int) bool {
		return t[i] > t[j]
	})
	if strings.Compare(string(t), string(s)) > 0 {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}
