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
	h, _ := strconv.Atoi(s[0])
	a, _ := strconv.Atoi(s[1])
	if h%a != 0 {
		fmt.Fprint(writer, (h/a)+1)
	} else {
		fmt.Fprint(writer, h/a)
	}

}
