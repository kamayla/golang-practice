package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("なにこの記法")
		}
	}(file)

	_, err2 := file.Write([]byte("Hello World2"))

	if err2 != nil {
		return
	}
}
