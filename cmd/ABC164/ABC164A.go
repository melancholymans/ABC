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
	s, _ := strconv.Atoi(r[0])
	w, _ := strconv.Atoi(r[1])
	if w >= s {
		fmt.Fprintln(writer, "unsafe")
	} else {
		fmt.Fprintln(writer, "safe")
	}
}
