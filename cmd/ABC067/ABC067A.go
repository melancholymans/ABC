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
	a, _ := strconv.Atoi(r[0])
	b, _ := strconv.Atoi(r[1])
	if a%3 == 0 {
		fmt.Fprintln(writer, "Possible")
		return
	} else if b%3 == 0 {
		fmt.Fprintln(writer, "Possible")
		return
	} else if (a+b)%3 == 0 {
		fmt.Fprintln(writer, "Possible")
		return
	}
	fmt.Fprintln(writer, "Impossible")
}
