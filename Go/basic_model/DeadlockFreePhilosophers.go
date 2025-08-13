package main

import "fmt"

func philosopher(id int, left, right chan struct{}) {
	for {
		fmt.Println("PHIL", id, ": sitdown")
		if id%2 == 0 {
			<-left
			<-right
		} else {
			<-right
			<-left
		}
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
		left := forks[(i-1+N)%N]
		right := forks[i]
		go philosopher(i, left, right)
	}
	select {}
}
