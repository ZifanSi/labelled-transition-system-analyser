package main

import "fmt"

func resource(acq, rel chan struct{}) {
	for {
		<-acq
		<-rel
	}
}

func user(acq, rel chan struct{}, name string) {
	for {
		fmt.Println(name, "acquire")
		acq <- struct{}{}
		fmt.Println(name, "use")
		fmt.Println(name, "release")
		rel <- struct{}{}
	}
}

func main() {
	acq := make(chan struct{})
	rel := make(chan struct{})

	go resource(acq, rel)
	go user(acq, rel, "A")
	go user(acq, rel, "B")

	select {}
}
