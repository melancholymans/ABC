package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	s := sc.Text()
	s1, _ := strconv.Atoi(string(s[0]))
	s2, _ := strconv.Atoi(string(s[1]))
	s3, _ := strconv.Atoi(string(s[2]))
	fmt.Fprintln(writer, s1+s2+s3)
}
