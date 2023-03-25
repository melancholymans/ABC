package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Mountain struct {
	name   string
	height int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	ad := make([]Mountain, 0)
	for i := 0; i < n; i++ {
		sc.Scan()
		a := strings.Split(sc.Text(), " ")
		s := a[0]
		t, _ := strconv.Atoi(a[1])
		var m Mountain
		m.name = s
		m.height = t
		ad = append(ad, m)
	}
	sort.Slice(ad, func(i, j int) bool {
		return ad[i].height > ad[j].height
	})
	fmt.Fprintln(writer, ad[1].name)
}
