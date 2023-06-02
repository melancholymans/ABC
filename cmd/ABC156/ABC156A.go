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
	n, _ := strconv.Atoi(r1[0])
	r, _ := strconv.Atoi(r1[1])
	if n >= 10 {
		fmt.Fprintln(writer, r)
	} else {
		fmt.Fprintln(writer, r+100*(10-n))
	}
}
