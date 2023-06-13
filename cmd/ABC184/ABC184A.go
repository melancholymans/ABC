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
	a, _ := strconv.Atoi(r1[0])
	b, _ := strconv.Atoi(r1[1])
	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	c, _ := strconv.Atoi(r2[0])
	d, _ := strconv.Atoi(r2[1])
	fmt.Fprintln(writer, a*d-b*c)
}
