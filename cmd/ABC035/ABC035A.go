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
	w, _ := strconv.Atoi(r[0])
	h, _ := strconv.Atoi(r[1])
	fmt.Fprintln(writer, w, h)
}
