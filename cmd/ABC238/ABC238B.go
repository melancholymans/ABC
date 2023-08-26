package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := ioutil.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}
}
func ni() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
func nis(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ni()
	}
	return a
}
func ni2() (int, int) {
	return ni(), ni()
}
func ni3() (int, int, int) {
	return ni(), ni(), ni()
}
func ni4() (int, int, int, int) {
	return ni(), ni(), ni(), ni()
}
func ns() string {
	sc.Scan()
	return sc.Text()
}
func nf() float64 {
	sc.Scan()
	f, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return f
}

func main() {
	defer wtr.Flush()
	n := ni()
	sl := make([]bool, 360)
	sl[0] = true
	cum := 0
	for i := 0; i < n; i++ {
		cum += ni()
		cum = cum % 360
		sl[cum] = true
	}
	r, c := 0, 0
	for i := 0; i <= 360; i++ {
		if sl[i%360] {
			r = Max(r, c)
			c = 0
		}
		c += 1
	}
	fmt.Fprintln(wtr, r)
}

/*
	int main(){
	  int p=0;
	  for(int i=0;i<n;i++){
	    int a;
	    cin >> a;
	    p+=a;p%=360;
	    fl[p]=true;
	  }
	  int res=0,cur=0;
	  for(int i=0;i<=360;i++){
	    if(fl[i%360]){
	      res=max(res,cur);
	      cur=0;
	    }
	    cur++;
	  }
	  cout << res << '\n';
	}
*/
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
