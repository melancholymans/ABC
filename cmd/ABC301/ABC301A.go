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
	//n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	s := sc.Text()
	CT := 0
	CA := 0
	last := ""
	for _, v := range s {
		if v == 'T' {
			CT += 1
			last = "T"
		}
		if v == 'A' {
			CA += 1
			last = "A"
		}
	}
	if CT > CA {
		fmt.Fprintln(writer, "T")
	} else if CT == CA {
		if last == "A" {
			fmt.Fprintln(writer, "T")
		} else {
			fmt.Fprintln(writer, "A")
		}
	} else {
		fmt.Fprintln(writer, "A")
	}
}
