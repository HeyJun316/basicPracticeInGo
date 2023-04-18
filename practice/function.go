package main

import (
	"time"
	"fmt"
)

// func Plus(x, y int) int {
// 	return x + y
// }

// func Div(x, y int) (int, int) {
// 	q := x/y
// 	r := x % y
// 	return q, r
// }

// // 引数と返り値がわかる
// func D(price int) (result int) {
// 	result = price * 2
// 	return
// }

func ReturnFunc() func() {
	return func() {
		fmt.Println("im")
	}
}

// func CallFunction(f func()) {
// 	f()
// }
// クロージャ
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

// ジェネレータ
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// 片アサーション
func anything(a interface{}) {
	fmt.Println(a)

	var x interface{} = 3
	i := x.(int)
	fmt.Println(i +2)

}
func sub() {
	for{
		fmt.Println("sub loop")
		time.Sleep(100 * time.Millisecond)
	}
}
func main() {
	//並行処理
	// init() mainより早く行える初期化処理
	go sub()
	go sub()

	for {
	fmt.Println("mainloop")
	time.Sleep(200 * time.Millisecond)
	}

	/*
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
		}()
		panic("runtime error")
		fmt.Println("start")
		*/
	// Loop:
	// 	for {
		// 		for{
	// 			for{
	// 				fmt.Println("START")
	// 				break
	// 			}
	// 			fmt.Println("PRocessing")
	// 		}
	// 		fmt.Println("END")
	// 	}
	// Loop:
	// 	for i:=0;i <3;i++ {
	// 		for j := 1; j <3 ;j++ {
	// 			if j > 1 {
	// 				continue Loop
	// 			}
	// 			fmt.Println(i,j,i*j)
	// 		}
	// 		fmt.Println("processing")
	// 	}
	/*
	file, err := os.Create("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Write([]byte("hello"))
	*/
	// anything("aa")
	// anything(1)
	/*
	i:=0
	for {
		i++
		if i == 3 {
			break
		}
		fmt.Println("aa")
	}


	if b:= 100; b ==100 {
		fmt.Println("one")
	}
	ints := integers()
	fmt.Println(ints())

	f := Later()
	fmt.Println(f("a"))
	fmt.Println(f("asd"))
	fmt.Println(f("asdf"))
	fmt.Println(f("aadf"))


	arr := [3]int{1,2,3}
	for i := 0; i<len(arr) ;i++{
		fmt.Println(arr[i])
	}
	*/

	// indexとvalueを取得している
	// for i,v := range arr{
	// 	fmt.Println(i,v)
	// }

	// sl := []string{"pyoun", "php", "go"}
	// for i,v := range sl {
	// 	fmt.Println(i,v)
	// }

	/*
	sm := map[string]int{"aap":100, "bana": 200}
	for k,v := range {
		fmt.Println(k)
	}
	*/
	// CallFunction(func() {
		// fmt.Println("i")
	// })
	// CallFunction()
	// a :=ReturnFunc ()
	// a()
	// i := Plus(1,2)
	// fmt.Println(i)

	// // i2, i3 := Div(9,4)
	// i2, _ := Div(9,4)
	// fmt.Println(i2)

	// f := func(j,k int) int {
	// 	return j+k
	// }

	// r := f(1,2)
	// fmt.Println(r)

	// // 渡す値を指定できる
	// d := func(j,k int) int {
	// 	return j+k
	// }(3,1)
	// fmt.Println(d)

}