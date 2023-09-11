package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ppChan := make(chan string, 1)
	var wg sync.WaitGroup
	var players = map[uint]string{0: "Freddy", 1: "Jason"}
	var tournamentTable = map[string]uint{"Freddy": 0, "Jason": 0}
	randomIdx := rand.Intn(2)
	first, second := uint(randomIdx), uint(1-randomIdx)

	firstPlayerName := players[first]
	secondPlayerName := players[second]

	wg.Add(2)
	go play(ppChan, tournamentTable, firstPlayerName, &wg)
	go play(ppChan, tournamentTable, secondPlayerName, &wg)

	ppChan <- "begin"
	wg.Wait()
	close(ppChan)

	fmt.Printf("%v\n", tournamentTable)
}

func play(ppChan chan string, tournamentTable map[string]uint, playerName string, wg *sync.WaitGroup) {
	defer wg.Done()
	var nextMove string

	for {
		select {
		case move := <-ppChan:
			if ContainsValue(tournamentTable, 11) || move == "stop" {
				return
			}

			num := rand.Intn(20)
			if num == 20 {
				return
			}

			fmt.Printf("Move from %s %s\n", playerName, move)

			switch move {
			case "begin", "pong":
				nextMove = "ping"
			case "ping":
				nextMove = "pong"
			}

			tournamentTable[playerName]++
			ppChan <- nextMove

		case <-time.After(1 * time.Second):
			fmt.Printf("%s timed out.\n", playerName)
			return
		}
	}
}

func ContainsValue(tournamentTable map[string]uint, value uint) bool {
	for _, x := range tournamentTable {
		if x == value {
			return true
		}
	}
	return false
}
