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
	r := strings.Split(sc.Text(), " ")
	d, _ := strconv.Atoi(r[0])
	x, _ := strconv.Atoi(r[1])
	total := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		for j := 0; ; j++ {
			if j == 0 {
				total += 1
				continue
			}
			if j*a+1 > d {
				break
			}
			total += 1
		}
	}
	fmt.Fprintln(writer, total+x)
}
