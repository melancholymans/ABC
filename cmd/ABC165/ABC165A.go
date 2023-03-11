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
	s := strings.Split(sc.Text(), " ")
	k, _ := strconv.Atoi(s[0])
	sc.Scan()
	s = strings.Split(sc.Text(), " ")
	a, _ := strconv.Atoi(s[0])
	b, _ := strconv.Atoi(s[1])

	for i := 1; i < (b/k)+1; i++ {
		if i*k >= a && i*k <= b {
			fmt.Fprint(writer, "OK")
			goto Z
		}
	}
	fmt.Fprint(writer, "NG")
Z:
}
