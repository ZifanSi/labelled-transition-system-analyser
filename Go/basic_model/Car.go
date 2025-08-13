package main

import (
	"fmt"
	"time"
)

const N = 4

func carpark(arrReq, depReq, arrEv, depEv chan struct{}) {
	free := N
	for {
		select {
		case <-arrReq:
			if free > 0 {
				free--
				arrEv <- struct{}{}
			}
		case <-depReq:
			if free < N {
				free++
				depEv <- struct{}{}
			}
		}
	}
}

func arrivals(ch chan struct{}, d time.Duration) { for { ch <- struct{}{}; time.Sleep(d) } }
func departures(ch chan struct{}, d time.Duration) { for { ch <- struct{}{}; time.Sleep(d) } }

func overflow(arrEv, depEv chan struct{}) {
	c := 0
	for {
		select {
		case <-arrEv:
			c++
			if c > N { fmt.Println("OVERFLOW"); return }
		case <-depEv:
			c--
			if c < 0 { fmt.Println("UNDERFLOW"); return }
		}
	}
}

func main() {
	arrReq, depReq := make(chan struct{}), make(chan struct{})
	arrEv, depEv := make(chan struct{}), make(chan struct{})

	go carpark(arrReq, depReq, arrEv, depEv)
	go arrivals(arrReq, 300*time.Millisecond)
	go departures(depReq, 500*time.Millisecond)
	go overflow(arrEv, depEv)

	select {}
}
