package main

import "fmt"

func client(call chan struct{}, reply chan struct{}) {
	for {
		fmt.Println("CLIENT: call")
		call <- struct{}{}
		<-reply
		fmt.Println("CLIENT: continue")
	}
}

func server(request chan struct{}, reply chan struct{}) {
	for {
		<-request
		fmt.Println("SERVER: service")
		reply <- struct{}{}
	}
}

func main() {
	call := make(chan struct{})
	reply := make(chan struct{})
	go client(call, reply)
	go server(call, reply)
	select {}
}
