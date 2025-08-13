package main

import "fmt"

func philosopher(id int, right, left chan struct{}) {
	for {
		fmt.Println("PHIL", id, ": sitdown")
		<-right
		<-left
		fmt.Println("PHIL", id, ": eat")
		left <- struct{}{}
		right <- struct{}{}
		fmt.Println("PHIL", id, ": arise")
	}
}

func fork() chan struct{} {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	return ch
}

func main() {
	const N = 5
	forks := make([]chan struct{}, N)
	for i := 0; i < N; i++ {
		forks[i] = fork()
	}
	for i := 0; i < N; i++ {
		right := forks[i]
		left := forks[(i-1+N)%N]
		go philosopher(i, right, left)
	}
	select {}
}
