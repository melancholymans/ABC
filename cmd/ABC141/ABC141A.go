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
	r := sc.Text()
	switch r {
	case "Sunny":
		fmt.Fprintln(writer, "Cloudy")
	case "Cloudy":
		fmt.Fprintln(writer, "Rainy")
	case "Rainy":
		fmt.Fprintln(writer, "Sunny")
	}
}
