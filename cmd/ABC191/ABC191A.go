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
	v, _ := strconv.Atoi(r[0])
	t, _ := strconv.Atoi(r[1])
	s, _ := strconv.Atoi(r[2])
	d, _ := strconv.Atoi(r[3])
	if v*t <= d && v*s >= d {
		fmt.Fprintln(writer, "No")
	} else {
		fmt.Fprintln(writer, "Yes")
	}
}
