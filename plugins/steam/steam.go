package main

import (
	"fmt"

	"github.com/Stumblinbear/gridlock/api"
)

var Name string = "Steam Sync"
var Version string = "0.0.1"
var Author string = "Gridlock"
var Description string = "Load your steam library"

func main() {
	fmt.Println(Name, "version", Version)
}

func OnStart(g *api.API) error {
	fmt.Println("Steam OnStart called")

	g.QueueNotification(api.Notification{
		Title: "Steam",
		Text: "Syncing library...",
	})

	return nil
}
