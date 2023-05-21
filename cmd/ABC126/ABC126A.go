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
	//n, _ := strconv.Atoi(r[0])
	k, _ := strconv.Atoi(r[1])
	sc.Scan()
	s := []byte(sc.Text())
	if s[k-1] < 68 {
		s[k-1] += 32
	}
	fmt.Fprintln(writer, string(s))
}
