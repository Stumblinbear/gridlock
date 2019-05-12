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

	// Only add this launcher if parsec is installed
	gapi.AddLauncher("parsec", api.Launcher{
		Name: "Parsec",
		Description: "Launch game and stream it through Parsec",
		
		Require: true,

		CanStart: func(gq api.GameQuery) bool {
			// Ensure parsec exists on the current system
			// If this launcher was displayed for a remote system, we already know that it has it installed
			return true
		},
		StartGame: func(gq api.GameQuery) error {
			// If parsec is not started, start it
			// If query to a remote system, somehow, we should forward the request to the remote system. It should be a blocking call.
			// Here, we should start up the parsec connection.
			return nil
		},
	})

	return nil
}
