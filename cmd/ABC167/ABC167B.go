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
	//c, _ := strconv.Atoi(r[2])
	k, _ := strconv.Atoi(r[3])
	if a >= k {
		fmt.Fprintln(writer, k)
	} else if a < k && (a+b) >= k {
		fmt.Fprintln(writer, a)
	} else {
		fmt.Fprintln(writer, a+((a+b)-k))
	}
}