package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	s := sc.Text()
	mp := map[rune]int{}
	for _, v := range s {
		mp[v] += 1
	}
	if len(mp) != 2 {
		fmt.Fprintln(writer, "No")
		return
	}
	for _, v := range mp {
		if v != 2 {
			fmt.Fprintln(writer, "No")
			return
		}
	}
	fmt.Fprintln(writer, "Yes")
}
