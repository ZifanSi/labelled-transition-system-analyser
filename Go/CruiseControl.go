package main

import (
	"fmt"
)

func cruiseController(events <-chan string, out chan<- string) {
	state := "INACTIVE"
	for e := range events {
		switch state {
		case "INACTIVE":
			if e == "engineOn" {
				out <- "clearSpeed"
				state = "ACTIVE"
			}
		case "ACTIVE":
			switch e {
			case "engineOff":
				state = "INACTIVE"
			case "on":
				out <- "recordSpeed"
				out <- "enableControl"
				state = "CRUISING"
			case "off", "brake", "accelerator":
				// remain ACTIVE
			}
		case "CRUISING":
			switch e {
			case "engineOff":
				state = "INACTIVE"
			case "off", "brake", "accelerator":
				out <- "disableControl"
				state = "STANDBY"
			case "on":
				out <- "recordSpeed"
				out <- "enableControl"
			}
		case "STANDBY":
			switch e {
			case "engineOff":
				state = "INACTIVE"
			case "resume":
				out <- "enableControl"
				state = "CRUISING"
			case "on":
				out <- "recordSpeed"
				out <- "enableControl"
				state = "CRUISING"
			}
		}
		fmt.Println("CRUISECONTROLLER state =", state, "after event", e)
	}
}

func speedControl(events <-chan string) {
	state := "DISABLED"
	for e := range events {
		switch state {
		case "DISABLED":
			if e == "enableControl" {
				state = "ENABLED"
			}
		case "ENABLED":
			switch e {
			case "speed":
				fmt.Println("setThrottle")
			case "disableControl":
				state = "DISABLED"
			}
		}
		fmt.Println("SPEEDCONTROL state =", state, "after event", e)
	}
}

func main() {
	events := make(chan string)
	ctrlOut := make(chan string)

	go cruiseController(events, ctrlOut)
	go speedControl(ctrlOut)

	go func() {
		events <- "engineOn"
		events <- "on"
		events <- "speed"
		events <- "brake"
		events <- "resume"
		close(events)
	}()

	select {}
}
