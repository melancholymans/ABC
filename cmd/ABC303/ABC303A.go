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
	//n := sc.Text()
	sc.Scan()
	s := sc.Text()
	sc.Scan()
	t := sc.Text()
	s = strings.Replace(s, "l", "1", -1)
	s = strings.Replace(s, "o", "0", -1)
	t = strings.Replace(t, "l", "1", -1)
	t = strings.Replace(t, "o", "0", -1)
	if strings.Compare(s, t) == 0 {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}
