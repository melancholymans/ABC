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
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	s := []byte(sc.Text())
	ad := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i]%byte(91-n) >= 65 {
			ad = append(ad, s[i]+byte(n))
		} else {
			ad = append(ad, s[i]%byte(91-n)+65)
		}
	}
	fmt.Fprintln(writer, string(ad))
}
