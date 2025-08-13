package main

import (
	"fmt"
	"time"
)

type tupleSpace struct {
	count int
	max   int
}

func (t *tupleSpace) out(val string) {
	if t.count < t.max {
		t.count++
		fmt.Println("out", val, "count", t.count)
	}
}

func (t *tupleSpace) in(val string) {
	if t.count > 0 {
		t.count--
		fmt.Println("in", val, "count", t.count)
	}
}

func (t *tupleSpace) inp(val string) {
	if t.count > 0 {
		t.count--
		fmt.Println("inp true", val, "count", t.count)
	} else {
		fmt.Println("inp false", val, "count", t.count)
	}
}

func (t *tupleSpace) rd(val string) {
	fmt.Println("rd", val, "count", t.count)
}

func (t *tupleSpace) rdp(val string) {
	fmt.Println("rdp", val, "count", t.count)
}

func main() {
	ts := &tupleSpace{count: 0, max: 2}
	go func() {
		ts.out("any")
		ts.out("any")
		ts.rd("any")
		ts.in("any")
		ts.inp("any")
		ts.rdp("any")
	}()
	time.Sleep(time.Second)
}
