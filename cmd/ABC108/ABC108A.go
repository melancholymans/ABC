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
	k, _ := strconv.Atoi(sc.Text())
	if k%2 == 0 {
		fmt.Fprintln(writer, k/2*k/2)
	} else {
		fmt.Fprintln(writer, (k/2+1)*(k/2))
	}
}
