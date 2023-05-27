package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	n, _ := strconv.ParseFloat(sc.Text(), 64)
	if math.Mod(n, 2) == 0 {
		fmt.Fprintln(writer, (n/2)/n)
	} else {
		fmt.Fprintln(writer, (math.Floor(n/2)+1)/n)
	}
}
