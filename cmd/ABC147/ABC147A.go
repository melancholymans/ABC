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
	a1, _ := strconv.Atoi(r[0])
	a2, _ := strconv.Atoi(r[1])
	a3, _ := strconv.Atoi(r[2])
	if a1+a2+a3 >= 22 {
		fmt.Fprintln(writer, "bust")
	} else {
		fmt.Fprintln(writer, "win")
	}
}
