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
	s := []byte(sc.Text())
	l := len(s)
	f := false
	f1 := true
	count := 0
	if s[0] != 65 { //65=='A'
		f = true
		goto Z
	}
	for i := 1; i < l; i++ { //A Cだけ大文字であること
		c := s[i]
		if c < 97 && c != 67 { //大文字であればアウト、ただし大文字のCはOK
			f = true
			goto Z
		}
		if c == 67 {
			count += 1
			if count > 1 {
				f = true
				goto Z
			}
		}
	}
	for i := 2; i < l-1; i++ {
		if s[i] == 67 {
			f1 = false
		}
	}
	if f1 {
		f = true
	}
Z:
	if f {
		fmt.Fprintln(writer, "WA")
	} else {
		fmt.Fprintln(writer, "AC")
	}
}
