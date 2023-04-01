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
	s1 := sc.Text()
	sc.Scan()
	s2 := sc.Text()
	sc.Scan()
	s3 := sc.Text()
	sc.Scan()
	s4 := sc.Text()
	calc(s1, s2, s3, s4)
}

func calc(s1, s2, s3, s4 string) {
	h := []string{"H", "2B", "3B", "HR"}
	count := 0
	for _, s := range h {
		if s1 == s {
			count += 1
			continue
		}
		if s2 == s {
			count += 1
			continue
		}
		if s3 == s {
			count += 1
			continue
		}
		if s4 == s {
			count += 1
			continue
		}
	}
	if count == 4 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
