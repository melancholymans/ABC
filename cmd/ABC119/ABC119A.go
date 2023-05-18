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
	r := strings.Split(sc.Text(), "/")
	//y:=r[0]
	m, _ := strconv.Atoi(r[1])
	//d:=r[2]
	if m > 4 {
		fmt.Fprintln(writer, "TBD")
	} else {
		fmt.Fprintln(writer, "Heisei")
	}
}
