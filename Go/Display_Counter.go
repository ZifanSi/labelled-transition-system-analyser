package main

import "fmt"

const N = 4

func counter(inc chan struct{}, display chan struct{}) {
	v := 0
	for {
		select {
		case <-inc:
			if v < N {
				v++
			}
		case <-display:
			fmt.Println(v)
		}
	}
}

func main() {
	inc := make(chan struct{})
	display := make(chan struct{})

	go counter(inc, display)

	go func() {
		inc <- struct{}{}
		inc <- struct{}{}
		display <- struct{}{}
		inc <- struct{}{}
		display <- struct{}{}
	}()

	select {}
}
