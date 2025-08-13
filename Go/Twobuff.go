package main

import "fmt"

func buff(in chan int, out chan int, name string) {
	for v := range in {
		fmt.Println(name, "in", v)
		out <- v
		fmt.Println(name, "out", v)
	}
	close(out)
}

func main() {
	aIn := make(chan int)
	aOut := make(chan int)
	bOut := make(chan int)

	go buff(aIn, aOut, "A")
	go buff(aOut, bOut, "B")

	go func() {
		for i := 0; i <= 3; i++ {
			fmt.Println("IN", i)
			aIn <- i
		}
		close(aIn)
	}()

	for v := range bOut {
		fmt.Println("OUT", v)
	}
}
