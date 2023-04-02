package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	s1 := strings.Split(sc.Text(), " ")
	Reversei(s1)
	sc.Scan()
	s2 := strings.Split(sc.Text(), " ")
	Reversei(s2)
	sc.Scan()
	s3 := strings.Split(sc.Text(), " ")
	Reversei(s3)
	sc.Scan()
	s4 := strings.Split(sc.Text(), " ")
	Reversei(s4)
	fmt.Fprintln(writer, strings.Join(s4, " "))
	fmt.Fprintln(writer, strings.Join(s3, " "))
	fmt.Fprintln(writer, strings.Join(s2, " "))
	fmt.Fprintln(writer, strings.Join(s1, " "))
}

func Reversei(sl []string) {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
}
