package main

import (
	"fmt"
)

func loop(id int, mutex chan struct{}) {
	for {
		<-mutex
		fmt.Println("p", id, ": enter")
		fmt.Println("p", id, ": exit")
		mutex <- struct{}{}
	}
}

func main() {
	mutex := make(chan struct{}, 1)
	mutex <- struct{}{}

	for i := 1; i <= 3; i++ {
		go loop(i, mutex)
	}

	select {}
}
