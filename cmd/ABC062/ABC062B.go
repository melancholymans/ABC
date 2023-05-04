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
	h, _ := strconv.Atoi(r[0])
	w, _ := strconv.Atoi(r[1])
	for i := 0; i < h+2; i++ {
		if i == 0 || i == h+1 {
			fmt.Fprintln(writer, strings.Repeat("#", w+2))
		} else {
			sc.Scan()
			a := "#" + sc.Text() + "#"
			fmt.Fprintln(writer, a)
		}
	}
}
