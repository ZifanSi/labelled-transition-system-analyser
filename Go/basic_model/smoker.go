package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Ingredient string

const (
	Tobacco Ingredient = "Tobacco"
	Paper   Ingredient = "Paper"
	Match   Ingredient = "Match"
)

func agent(table chan [2]Ingredient) {
	ingredients := [3]Ingredient{Tobacco, Paper, Match}
	for {
		i, j := rand.Intn(3), rand.Intn(3)
		if i == j {
			continue
		}
		table <- [2]Ingredient{ingredients[i], ingredients[j]}
	}
}

func smoker(name string, own Ingredient, table chan [2]Ingredient) {
	for {
		ing := <-table
		if ing[0] != own && ing[1] != own {
			fmt.Printf("%s picks up %s and %s, rolls cigarette, smokes\n", name, ing[0], ing[1])
			time.Sleep(time.Millisecond * 200)
		} else {
			table <- ing // put back if not yours
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	table := make(chan [2]Ingredient, 1)

	go agent(table)
	go smoker("Smoker T", Tobacco, table)
	go smoker("Smoker P", Paper, table)
	go smoker("Smoker M", Match, table)

	select {}
}
