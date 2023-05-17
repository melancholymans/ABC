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
	ab, _ := strconv.Atoi(r[0])
	bc, _ := strconv.Atoi(r[1])
	//ca, _ := strconv.Atoi(r[2])
	fmt.Fprintln(writer, ab*bc/2)
}
