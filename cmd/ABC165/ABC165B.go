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
	x, _ := strconv.Atoi(sc.Text())
	m := float64(100)
	for i := 1; ; i++ {
		m = m + float64(int(m*0.01))
		if x <= int(m) {
			fmt.Fprintln(writer, i)
			break
		}
	}
}
