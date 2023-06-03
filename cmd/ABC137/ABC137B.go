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
	k, _ := strconv.Atoi(r[0])
	x, _ := strconv.Atoi(r[1])
	for i := x - (k - 1); i < x+k; i++ {
		fmt.Fprint(writer, i, " ")
	}
	fmt.Fprint(writer, "")

}
