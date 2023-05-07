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
	r := sc.Text()
	n, _ := strconv.Atoi(r)
	tn := n
	total := 0
	fmt.Println("start", tn, total)
	for i := 0; i < len(r); i++ {
		total += tn % 10
		tn /= 10
		fmt.Println("i=", i, tn, total)
	}
	fmt.Println("end", tn, total)
	if n%total == 0 {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}
