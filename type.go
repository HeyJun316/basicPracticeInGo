package main

import (
	"fmt"
)
// import (
// 	"fmt"
// 	"strconv"
// )

// 定数 文頭を大文字で全てのパッケージから、小文字の場合このパッケージのみから定数を呼び出せる
const Pi = 3.14
const (
	URL = "http"
	SiteName = "test"
)

const (
	A = 1
	B //1
	D //1
	f = 4
)

// iota連続した数字を自動で生成
const (
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println(c0,c1,c2)

	fmt.Println(Pi)
	// Pi = 1
	var i int = 100
	/*
	intは環境依存で決められる
	int8/16/32/64
	*/
  fmt.Println(URL, SiteName)
  fmt.Println(A,B,D,f)

	var i2 int64 = 200
	fmt.Println(i + 50)
	// fmt.Println(i + i2)
	// Printfと%T 型を調べる
	fmt.Printf("%T\n", i2)
	// 型変換
	fmt.Printf("%T\n", int32(i2))

	fmt.Println(i+int(i2))

	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n",fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	pinf := 1.0/ zero
	fmt.Println(pinf)

	ninf := -1.0/ zero
	fmt.Println(ninf)

	nan := zero/ zero
	fmt.Println(nan)

	// var t,f bool = true, false
	var s string = "hello world"
	fmt.Println(string(s[1]))

	byteA := []byte{72, 73}
	fmt.Println(byteA)

	fmt.Printf(string(byteA))

	c := []byte("hi")
	fmt.Println(c)

	// 配列 要素数を追加できない
	// スライスで変更できるらしい
	// 配列型＝要素数を変更できない。

	// スライス型＝要素数を変更可能。
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)


	arr2 := [3]int{1,2,3}
	fmt.Println(arr2)

	// 与えられて要素数が、[]に入る
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	var arr3 [3]string = [3]string{"A","N"}
	fmt.Println(arr3)
	// 代入
	fmt.Println(arr2[0])
	arr3[2] ="c"

	fmt.Println(arr3)

	// arr5 = arr1
	// fmt.Println(arr5)

	fmt.Println(len(arr1))

 // nilという特殊な形を持つ
 // interface
 // あらゆる型を代入できる
 // 演算の対象ではない
  var x interface{}
  fmt.Println(x)

  x = 3.14
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// 型変換
	// var k int =1
	// fl64k := float64(k)
	// fmt.Printf("%T\n", k)
	// fmt.Printf("%T\n", fl64k)
  // fmt.Println(fl64k)

	// fl := 10.5
	// var ss string = "100"
	// fmt.Printf("%T\n", ss)

	// _で返ってくる値を破棄
	// int, _ := strconv.Atoi(ss)
	// int, err := strconv.Atoi(ss)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("%T\n", int)
 /*
 var k2 int = 200
 k3 := strconv.Itoa(k2)
 fmt.Printf("%T\n", k3)
 fmt.Println(k3)
 */

	var h string = "hello"
	b := []byte(h)
	fmt.Println(b)


}