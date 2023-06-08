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
	for i := 0; i < len(r); i++ {
		x, _ := strconv.Atoi(r[i])
		if x == 0 {
			fmt.Fprintln(writer, i+1)
		}
	}
}
