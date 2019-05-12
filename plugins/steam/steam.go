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

func OnStart(gapi *api.API) error {
	fmt.Println("Steam OnStart called")

	gapi.AddPlatform(api.Platform{
		Name: "Steam",
	})

	//gapi.AddGame()

	gapi.AddMetadataResolver(api.MetadataResolver{
		Name: "Steam",
		Resolve: func(name string) api.Metadata {
			return api.Metadata{}
		},
	})

	go gapi.QueueNotification(api.Notification{
		Title: "Steam",
		Text: "Syncing library...",
	})

	return nil
}
