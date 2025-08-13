package main

import "fmt"

func maze(start int, moves chan string) {
	transitions := map[int]map[string]int{
		0: {"north": -1, "east": 1},
		1: {"east": 2, "south": 4, "west": 0},
		2: {"south": 5, "west": 1},
		3: {"east": 4, "south": 6},
		4: {"north": 1, "west": 3},
		5: {"north": 2, "south": 8},
		6: {"north": 3},
		7: {"east": 8},
		8: {"north": 5, "west": 7},
	}

	pos := start
	for m := range moves {
		next, ok := transitions[pos][m]
		if !ok {
			fmt.Println("Invalid move from position", pos)
			continue
		}
		if next == -1 {
			fmt.Println("STOP")
			return
		}
		pos = next
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
