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
	a, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	//b, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	//c, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	//d, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	e, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	k, _ := strconv.Atoi(sc.Text())
	if e-a > k {
		fmt.Fprintln(writer, ":(")
	} else {
		fmt.Fprintln(writer, "Yay!")
	}
}
