package main

import (
	ag "blood/ServerAgent"
	"log"
	"net/http"
)

func main() {

	var agent ag.ServerAgent = ag.ServerAgentInit()

	http.HandleFunc("/player", agent.GetPlayer)
	http.HandleFunc("/players", agent.GetPlayers)
	http.HandleFunc("/levelup", agent.LevelUp)
	http.HandleFunc("/getfile", agent.ServeFile)
	http.HandleFunc("/getLevel", agent.GetLevel)

	if err := http.ListenAndServe("localhost:5051", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
