package main

import "fmt"

func slave(start <-chan struct{}, slaveRotate chan<- struct{}, join <-chan struct{}) {
	for {
		<-start
		fmt.Println("Slave: start and rotate")
		slaveRotate <- struct{}{}
		<-join
		fmt.Println("Slave: joined back")
	}
}

func master(start chan<- struct{}, slaveRotate <-chan struct{}, join chan<- struct{}) {
	for {
		fmt.Println("Master: start slave")
		start <- struct{}{}
		<-slaveRotate
		fmt.Println("Master: rotate")
		join <- struct{}{}
	}
}

func main() {
	start := make(chan struct{})
	slaveRotate := make(chan struct{})
	join := make(chan struct{})
	go slave(start, slaveRotate, join)
	go master(start, slaveRotate, join)
	select {}
}
