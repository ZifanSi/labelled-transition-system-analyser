package main

import "fmt"

func port(send <-chan int, receive chan<- int) {
	queue := []int{}
	for {
		select {
		case msg := <-send:
			queue = append(queue, msg)
			fmt.Println("send", msg)
		default:
			if len(queue) > 0 {
				receive <- queue[0]
				fmt.Println("receive", queue[0])
				queue = queue[1:]
			}
		}
	}
}

func main() {
	send := make(chan int)
	receive := make(chan int)
	go port(send, receive)
	go func() {
		for i := 0; i < 10; i++ {
			send <- i
		}
	}()
	for {
		fmt.Println(<-receive)
	}
}
