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
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	w := strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		if calc(writer, w) {
			break
		} else {
			break
		}
	}
}

func calc(writer *bufio.Writer, sl []string) bool {
	wl := [5]string{"and", "not", "that", "the", "you"}
	for _, w := range wl {
		for _, s := range sl {
			if w == s {
				fmt.Fprintln(writer, "Yes")
				return true
			}
		}
	}
	fmt.Fprintln(writer, "No")
	return false
}
