package main

import (
	"fmt"
)

func main() {
	var nilMap map[string]int

	//varを使ったmapの宣言、ただし値の代入ができないのでこのような宣言はしない
	fmt.Println(len(nilMap))
	fmt.Println(nilMap["abc"])
	//nilMap["abc"] = 3

	//マップリテラルを使った宣言、int{}をつけることによって空のマテリアルで初期化していることになるので
	//値の代入も可能になる。
	totalWins := map[string]int{}
	totalWins["abc"] = 3
	fmt.Println(totalWins["abc"])

	//空でないマップマテリアルも設定可能
	teams := map[string][]string{
		"ライターズ": []string{"夏目", "森"},
		"ナイツ":   []string{"武田", "徳田"},
	}
	fmt.Println(teams)

	//makeを使った宣言の方法
	ages := make(map[string]int, 10)
	fmt.Println(ages)
	ages["GPT"] = 1000
	ages["Promt"] = 123
	fmt.Println(ages)

	//カンマ,OKイディオム
	//マップはまだ値が設定されていない場合はその変数のゼロ値を返す。そうすると設定してあったのが0を設定してあったのか、ゼロ値なのかがわからない問題が生じる
	//それを解決するのが「カンマ,OKイディオム」である。
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)
	v, ok = m["world"]
	fmt.Println(v, ok)
	v, ok = m["goodby"]
	fmt.Println(v, ok)
}
