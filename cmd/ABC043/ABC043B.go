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
	o := make([]rune, 0)
	for _, v := range s {
		switch v {
		case '0', '1':
			o = append(o, v)
		case 'B':
			if len(o) > 0 {
				o = append(o[:len(o)-1])
			}
		}
	}
	fmt.Fprintln(writer, string(o))
}
