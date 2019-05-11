package main

import (
	"fmt"

	"github.com/Stumblinbear/gridlock/api"
)

var Name string = "Parsec Launcher"
var Version string = "0.0.1"
var Author string = "Gridlock"
var Description string = "Stream via Parsec"

func main() {
	fmt.Println(Name, "version", Version)
}

func OnStart(gapi *api.API) error {
	fmt.Println("Parsec OnStart called")

	gapi.AddLauncher("parsec", api.Launcher{
		Name: "Parsec",
		CanStart: func(gq api.GameQuery) {
			return true
		},
		StartGame: func(gq api.GameQuery) {
			return nil
		},
	})

	return nil
}
