package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	s := []byte(sc.Text())
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			fmt.Fprintln(writer, "Bad")
			return
		}
	}
	fmt.Fprintln(writer, "Good")
}
