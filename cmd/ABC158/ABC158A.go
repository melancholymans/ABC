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
	s := sc.Text()
	if (s[0] == 65 && s[1] == 65 && s[2] == 65) || (s[0] == 66 && s[1] == 66 && s[2] == 66) {
		fmt.Fprintln(writer, "No")
		return
	}
	fmt.Fprintln(writer, "Yes")
}
