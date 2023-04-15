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
	sc.Scan()
	s := sc.Text()
	count := 0
	for i := 0; i < n; i++ {
		if s[i] == 'x' {
			fmt.Fprintln(writer, "No")
			return
		}
		if s[i] == 'o' {
			count += 1
		}
	}
	if count > 0 {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}
