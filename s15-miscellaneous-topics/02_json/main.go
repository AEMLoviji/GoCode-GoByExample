package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type SoccerClubs struct {
	Name    string
	Country string
	Value   float32 `json:"Dollar Value(B)"`
	Players []string
}

var teams = []SoccerClubs{
	{Name: "Manchester United", Country: "England", Value: 3.69,
		Players: []string{"Eric Cantona", "David Beckham", "Ryan Giggs"}},
	{Name: "Barcelona", Country: "Spain", Value: 3.64,
		Players: []string{"Johan Cruyff", "Xavi", "Andres Iniesta"}},
	{Name: "Real Madrid", Country: "Spain", Value: 3.58,
		Players: []string{"Zidane", "Roverto Carlos", "Ronaldo"}},
}

func main() {
	fmt.Println("\n- - - - - - - - - - - - - - - - - - - - - - - - - ")

	// use either json.Marshal() or json.MarshalIndent()
	//data, err := json.Marshal(teams)
	data, err := json.MarshalIndent(teams, "", "    ")

	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Printf("%v\n\n", teams)
	fmt.Printf("%+v\n\n", teams)
	fmt.Printf("%s\n", data)

	fmt.Println("\n- - - - - - - - - - - - - - - - - - - - - - - - - ")
	var names []struct{ Name string }

	if err := json.Unmarshal(data, &names); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	fmt.Println(names)
}
