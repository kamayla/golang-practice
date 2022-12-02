package main

// パッケージのインポート
import (
	"fmt"
	"time"
)

const PI = 3.14

const (
	URL      = "https://google.com"
	SiteName = "グーグル"
)

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println(PI)
	fmt.Println(URL)
	fmt.Println(SiteName)
	// fmtパッケージでConsoleに文字列を表示
	fmt.Println("hello world", time.Now())

	var count int = 100

	fmt.Println(count)

	var name string = "Ippei Kamimura"

	fmt.Println(name)

	var flagOne, flagTwo bool = false, true

	fmt.Println(flagOne, flagTwo)

	// 暗黙的定義
	// 型指定の必要がない
	a := 1234

	fmt.Println(a)

	fmt.Printf("%T\n", name)

	fmt.Println(`あいうえお
	あいうえお
		あいうえお`)

	fmt.Println(string(byte(71)))

	// 配列
	var arr [4]int

	fmt.Println(arr)

	var arr2 [4]int = [4]int{1, 3, 4, 5}

	fmt.Println(arr2)

	ints := integers()

	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	intsTwo := integers()

	fmt.Println(intsTwo())
	fmt.Println(intsTwo())
	fmt.Println(intsTwo())
	fmt.Println(intsTwo())

}
