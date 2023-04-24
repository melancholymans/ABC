package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	r := strings.Split(sc.Text(), " ")
	s1 := r[0]
	s2 := r[1]
	s3 := r[2]
	fmt.Fprintf(writer, "%s%s%s\n", string(s1[0]-32), string(s2[0]-32), string(s3[0]-32))
}
