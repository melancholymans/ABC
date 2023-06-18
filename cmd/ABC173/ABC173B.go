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
	n, _ := strconv.Atoi(sc.Text())
	mp := map[string]int{"AC": 0, "WA": 0, "TLE": 0, "RE": 0}
	for i := 0; i < n; i++ {
		sc.Scan()
		c := sc.Text()
		mp[c] += 1
	}
	fmt.Fprintln(writer, "AC x ", mp["AC"])
	fmt.Fprintln(writer, "WA x ", mp["WA"])
	fmt.Fprintln(writer, "TLE x ", mp["TLE"])
	fmt.Fprintln(writer, "RE x ", mp["RE"])
}
