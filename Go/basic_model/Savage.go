package main

import (
	"fmt"
	"time"
)

const (
	M       = 3
	SAVAGES = 4
)

func cook(pot chan struct{}, refill chan struct{}) {
	for range refill {
		for i := 0; i < M; i++ {
			pot <- struct{}{}
		}
		fmt.Println("Cook refills pot")
	}
}

func savage(id int, pot chan struct{}, refill chan struct{}) {
	for {
		select {
		case <-pot:
			fmt.Printf("Savage %d eats\n", id)
		default:
			select { case refill <- struct{}{}: default: }
			<-pot
			fmt.Printf("Savage %d eats\n", id)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	pot := make(chan struct{}, M)
	refill := make(chan struct{}, 1)
	for i := 0; i < M; i++ { pot <- struct{}{} }
	go cook(pot, refill)
	for i := 1; i <= SAVAGES; i++ {
		go savage(i, pot, refill)
	}
	select {}
}
