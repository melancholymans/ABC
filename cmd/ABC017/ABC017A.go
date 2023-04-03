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
	total := 0
	for i := 0; i < 3; i++ {
		sc.Scan()
		l := strings.Split(sc.Text(), " ")
		s, _ := strconv.Atoi(l[0])
		e, _ := strconv.Atoi(l[1])
		total += s * e / 10
	}
	fmt.Fprintln(writer, total)
}
