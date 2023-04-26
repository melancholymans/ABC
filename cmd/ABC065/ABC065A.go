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
	x, _ := strconv.Atoi(r[0])
	a, _ := strconv.Atoi(r[1])
	b, _ := strconv.Atoi(r[2])
	if -a+b <= 0 {
		fmt.Fprintln(writer, "delicious")
	} else if -a+b <= x {
		fmt.Fprintln(writer, "safe")
	} else {
		fmt.Fprintln(writer, "dangerous")
	}
}
