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
	n, _ := strconv.Atoi(r[0])
	k, _ := strconv.Atoi(r[1])
	if n%k == 0 {
		fmt.Fprintln(writer, 0)
	} else {
		fmt.Fprintln(writer, 1)
	}
}
