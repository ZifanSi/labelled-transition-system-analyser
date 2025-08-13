package main

import "fmt"

func maze(start int, moves chan string) {
	pos := start
	for m := range moves {
		switch pos {
		case 0:
			if m == "east" {
				pos = 1
			} else if m == "north" {
				fmt.Println("STOP")
				return
			}
		case 1:
			if m == "east" {
				pos = 2
			} else if m == "south" {
				pos = 4
			} else if m == "west" {
				pos = 0
			}
		case 2:
			if m == "south" {
				pos = 5
			} else if m == "west" {
				pos = 1
			}
		case 3:
			if m == "east" {
				pos = 4
			} else if m == "south" {
				pos = 6
			}
		case 4:
			if m == "north" {
				pos = 1
			} else if m == "west" {
				pos = 3
			}
		case 5:
			if m == "north" {
				pos = 2
			} else if m == "south" {
				pos = 8
			}
		case 6:
			if m == "north" {
				pos = 3
			}
		case 7:
			if m == "east" {
				pos = 8
			}
		case 8:
			if m == "north" {
				pos = 5
			} else if m == "west" {
				pos = 7
			}
		}
		fmt.Println("Moved", m, "-> at position", pos)
	}
}

func main() {
	moves := make(chan string)
	go maze(0, moves)
	go func() {
		moves <- "east"
		moves <- "south"
		moves <- "north"
		close(moves)
	}()
	select {}
}
