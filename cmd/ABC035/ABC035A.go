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
	if w%16 == 0 && h%9 == 0 {
		fmt.Fprintln(writer, "16:9")
	} else {
		fmt.Fprintln(writer, "4:3")
	}
}
