package main

import (
	"fmt"
	"time"
)

type ev struct{ kind, pat string }

func em(in <-chan ev, ls []chan ev) {
	for e := range in {
		if e.kind == "announce" {
			for _, l := range ls {
				l <- e
			}
		}
	}
}

func ann(out chan<- ev) {
	for _, p := range []string{"pat1", "pat2"} {
		out <- ev{"announce", p}
		time.Sleep(time.Second)
	}
}

func box(name, pat string, in <-chan ev) {
	for e := range in {
		if e.kind == "announce" && e.pat == pat {
			fmt.Println(name, "hit", pat)
			return
		}
	}
}

func main() {
	emc := make(chan ev)
	ls := []chan ev{{}, {}, {}, {}}
	for i := range ls {
		ls[i] = make(chan ev)
	}
	go em(emc, ls)
	go ann(emc)
	go box("a", "pat1", ls[0])
	go box("b", "pat2", ls[1])
	go box("c", "pat1", ls[2])
	go box("d", "pat2", ls[3])
	select {}
}
