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
	r1 := strings.Split(sc.Text(), " ")
	r, _ := strconv.Atoi(r1[0])
	d, _ := strconv.Atoi(r1[1])
	x, _ := strconv.Atoi(r1[2])
	for i := 0; i < 10; i++ {
		x = x*r - d
		fmt.Fprintln(writer, x)
	}
}
