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
	in := sc.Text()
	//i, j := 0, 0
	/*	for{
			o:=s[]
		}
	*/
	fmt.Fprintln(writer, in)
}
