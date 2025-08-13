package main

import "fmt"

const N = 5

func buffer(put <-chan struct{}, get <-chan struct{}) {
	count := 0
	for {
		select {
		case <-put:
			if count < N {
				count++
				fmt.Println("put, count =", count)
			}
		case <-get:
			if count > 0 {
				count--
				fmt.Println("get, count =", count)
			}
		}
	}
}

func producer(put chan<- struct{}) {
	for {
		put <- struct{}{}
	}
}

func consumer(get chan<- struct{}) {
	for {
		get <- struct{}{}
	}
}

func main() {
	put := make(chan struct{})
	get := make(chan struct{})

	go buffer(put, get)
	go producer(put)
	go consumer(get)

	select {}
}
