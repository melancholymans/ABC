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
	r := strings.Split(sc.Text(), " ")
	x1, _ := strconv.Atoi(r[0])
	y1, _ := strconv.Atoi(r[1])
	x2, _ := strconv.Atoi(r[2])
	y2, _ := strconv.Atoi(r[3])
	fmt.Fprintln(writer, x1, y1, x2, y2)
}
