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
	r := strings.Split(sc.Text(), " ")
	k, _ := strconv.Atoi(r[0])
	s, _ := strconv.Atoi(r[1])
	count := 0
	for x := 0; x <= k; x++ {
		for y := 0; y <= k; y++ {
			z := s - (x + y)
			if z >= 0 && z <= k {
				count += 1
			}
		}
	}
	fmt.Fprintln(writer, count)
}
