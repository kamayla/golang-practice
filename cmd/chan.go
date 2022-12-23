package main

import (
	"fmt"
	"time"
)

func reciever(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	ch := make(chan int, 2)

	go reciever("1.goroutin", ch)
	go reciever("2.goroutin", ch)
	go reciever("3.goroutin", ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}
	close(ch)
	time.Sleep(3 * time.Second)
}
