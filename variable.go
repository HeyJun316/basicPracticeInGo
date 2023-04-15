package main

import (
  "fmt"
)

// func main() {
// 	fmt.Println("hello world")
// 	fmt.Println(time.Now())
// }
	// i5 := 499
	var i5 int = 500
	// fmt.Println(i5)

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}
func main() {
	// 明示的な定義（基本的には型指定する）
	// var 変数名 肩=値
	var i int = 100
	fmt.Println(i)

	var s string = "hello Go"
	fmt.Println(s)

	var t,f bool = true, false
	fmt.Println(t,f)

	var (
		i2 int = 200
		s2 string = "Golang"
	)
	fmt.Println(i2,s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"

	i = 150
	fmt.Println(i)

	// 暗黙な定義
	// 変数名 := 値
	// 型推論で自動的に型入る
	// 関数ないでしか定義できない
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)
	fmt.Println(i5)

	// i4 := 500
	// fmt.Println(i4)
	// i4 = "string"
	// fmt.Println(i4)
	outer()
	// fmt.Println(s4)
	var s5 string = "Not use"
	fmt.Println(s5)
}