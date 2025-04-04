package main

import (
	"fmt"
	"go-starter/greetings"
)

type Player = greetings.Player
type Stack = greetings.Stack

func main() {
	players := Stack{"Team 1", []Player{}}
	players.AddPlayer(Player{"Andrew", 7, "Chicago Zephyrs"})
	players.AddPlayer(Player{"Ethan", 2, "Denver Broncos"})
	players.AddPlayer(Player{"Robert", 5, "Philadelphia 76ers"})
	players.AddPlayer(Player{"Ant", 76, "Phoenix Suns"})
	players.RemovePlayer(1)
	fmt.Println(players.List)

	var myMap map[string]uint8 = make(map[string]uint8) //unsigned ints are > 0
	var stars = map[string]string{"Bulls":"MJ", "Timberwolves":"KAT"} 
	myMap["David"] = 5
	fmt.Println(myMap["David"])
	fmt.Println(stars["Bulls"])
	delete(myMap, "David")

	var player1, inMap = stars["Pistons"]
	if inMap {
		fmt.Println(player1)
	} else {
		fmt.Println("Invalid Star")
	}

	for team, name := range stars {
		fmt.Printf("Team: %v, Name: %v \n", team, name)
	}

}