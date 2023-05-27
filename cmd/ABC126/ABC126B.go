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
	up, _ := strconv.Atoi(s[:2])
	do, _ := strconv.Atoi(s[2:])
	if (up >= 1 && up <= 12) && (do >= 1 && do <= 12) {
		fmt.Fprintln(writer, "AMBIGUOUS")
	} else {
		if up >= 1 && up <= 12 {
			fmt.Fprintln(writer, "MMYY")
		} else if do >= 1 && do <= 12 {
			fmt.Fprintln(writer, "YYMM")
		} else {
			fmt.Fprintln(writer, "NA")
		}
	}
}
