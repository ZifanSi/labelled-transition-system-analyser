package main

import "fmt"

const N = 4

func carpark(arrive <-chan struct{}, depart <-chan struct{}) {
	spaces := N
	for {
		select {
		case <-arrive:
			if spaces > 0 {
				spaces--
				fmt.Println("arrive, spaces =", spaces)
			}
		case <-depart:
			if spaces < N {
				spaces++
				fmt.Println("depart, spaces =", spaces)
			}
		}
	}
}

func main() {
	arrive := make(chan struct{})
	depart := make(chan struct{})

	go carpark(arrive, depart)

	go func() {
		arrive <- struct{}{}
		arrive <- struct{}{}
		depart <- struct{}{}
		arrive <- struct{}{}
	}()

	select {}
}
