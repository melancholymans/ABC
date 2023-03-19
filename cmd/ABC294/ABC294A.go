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
	a := strings.Split(sc.Text(), " ")
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		d, _ := strconv.Atoi(a[i])
		if d%2 == 0 {
			result = append(result, d)
		}
	}
	for i := 0; i < len(result); i++ {
		fmt.Fprint(writer, result[i], " ")
	}
	fmt.Fprint(writer, "\n")
}
