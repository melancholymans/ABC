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
	ad := []rune{}
	for i := 0; i < len(s); i++ {
		n, _ := strconv.Atoi(s[i])
		ad = append(ad, rune(n-1)+'a')
	}
	fmt.Println(string(ad))
}
