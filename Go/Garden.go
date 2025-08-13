package main

import "fmt"

const N = 4

func main() {
	v := 0
	eastArrive := make(chan struct{})
	westArrive := make(chan struct{})
	display := make(chan struct{})

	go func() {
		for {
			select {
			case <-eastArrive, <-westArrive:
				if v < N {
					v++
				}
			case <-display:
				fmt.Println("value", v)
			}
		}
	}()

	go func() {
		eastArrive <- struct{}{}
		westArrive <- struct{}{}
		display <- struct{}{}
	}()

	select {}
}
