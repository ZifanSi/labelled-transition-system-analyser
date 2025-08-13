package main

import "fmt"

func rotator(run, rotate, pause, interrupt chan struct{}, name string) {
	state := "PAUSED"
	for {
		switch state {
		case "PAUSED":
			select {
			case <-run:
				fmt.Println(name, "run")
				state = "RUN"
			case <-pause:
				fmt.Println(name, "pause")
			case <-interrupt:
				fmt.Println(name, "stop")
				return
			}
		case "RUN":
			select {
			case <-pause:
				fmt.Println(name, "pause")
				state = "PAUSED"
			case <-run:
				fmt.Println(name, "run")
			case <-rotate:
				fmt.Println(name, "rotate")
			case <-interrupt:
				fmt.Println(name, "stop")
				return
			}
		}
	}
}

func main() {
	aRun := make(chan struct{})
	aRotate := make(chan struct{})
	aPause := make(chan struct{})
	bRun := make(chan struct{})
	bRotate := make(chan struct{})
	bPause := make(chan struct{})
	stop := make(chan struct{})

	go rotator(aRun, aRotate, aPause, stop, "A")
	go rotator(bRun, bRotate, bPause, stop, "B")

	go func() {
		aRun <- struct{}{}
		aRotate <- struct{}{}
		bRun <- struct{}{}
		bPause <- struct{}{}
		stop <- struct{}{}
	}()

	select {}
}
