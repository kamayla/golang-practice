package main

import "fmt"

func main() {
	i := 1
	switch i {
	case 1:
		fmt.Println("これは1")
	case 2:
		fmt.Println("これは2")
	default:
		fmt.Println("ちゃんとしろ")
	}
}
