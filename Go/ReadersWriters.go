package main

import (
	"fmt"
	"sync"
	"time"
)

func reader(id int, rw *sync.RWMutex) {
	for {
		rw.RLock()
		fmt.Println("Reader", id, ": acquireRead")
		fmt.Println("Reader", id, ": examine")
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Reader", id, ": releaseRead")
		rw.RUnlock()
	}
}

func writer(id int, rw *sync.RWMutex) {
	for {
		rw.Lock()
		fmt.Println("Writer", id, ": acquireWrite")
		fmt.Println("Writer", id, ": modify")
		time.Sleep(300 * time.Millisecond)
		fmt.Println("Writer", id, ": releaseWrite")
		rw.Unlock()
	}
}

func main() {
	var rw sync.RWMutex
	Nread := 2
	Nwrite := 2

	for i := 1; i <= Nread; i++ {
		go reader(i, &rw)
	}
	for i := 1; i <= Nwrite; i++ {
		go writer(i, &rw)
	}

	select {}
}
