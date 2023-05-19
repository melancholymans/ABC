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
	H, _ := strconv.Atoi(r1[0])
	W, _ := strconv.Atoi(r1[1])
	sc.Scan()
	r2 := strings.Split(sc.Text(), " ")
	h, _ := strconv.Atoi(r2[0])
	w, _ := strconv.Atoi(r2[1])
	fmt.Fprintln(writer, H*W-h*W-w*H+w*h)
}
