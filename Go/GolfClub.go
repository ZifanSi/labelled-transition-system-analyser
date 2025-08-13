package main

import (
	"fmt"
	"math/rand"
	"time"
)

func allocator(get <-chan int, put <-chan int) {
	b := 5
	for {
		select {
		case n := <-get:
			if n <= b {
				b -= n
				fmt.Println("get", n, "left", b)
			}
		case n := <-put:
			b += n
			fmt.Println("put", n, "left", b)
		}
	}
}

func player(name string, need []int, get chan<- int, put chan<- int) {
	for {
		n := need[rand.Intn(len(need))]
		get <- n
		put <- n
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	get := make(chan int)
	put := make(chan int)
	go allocator(get, put)
	for _, p := range map[string][]int{"alice": {1, 2}, "bob": {1, 2}, "chris": {1, 2}, "dave": {3, 4, 5}, "eve": {3, 4, 5}} {
		go player("", p, get, put)
	}
	select {}
}
