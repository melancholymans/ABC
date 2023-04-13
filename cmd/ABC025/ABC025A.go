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
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	count := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			count += 1
			if count == n {
				fmt.Println(string(s[i]) + string(s[j]))
			}
		}
	}
}
