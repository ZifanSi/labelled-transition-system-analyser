package main

import (
	"fmt"
	"sync"
)

const N = 4

func main() {
	var v int
	var mu sync.Mutex
	eastArrive := make(chan struct{})
	westArrive := make(chan struct{})
	display := make(chan struct{})

	go func() {
		for {
			select {
			case <-eastArrive, <-westArrive:
				mu.Lock()
				if v < N {
					v++
				}
				mu.Unlock()
			case <-display:
				mu.Lock()
				fmt.Println("value", v)
				mu.Unlock()
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
