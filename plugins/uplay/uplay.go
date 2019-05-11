package main

import (
	"fmt"

	"github.com/Stumblinbear/gridlock/api"
)

var Name string = "UPlay Sync"
var Version string = "0.0.1"
var Author string = "Gridlock"
var Description string = "Load your UPlay library"

func main() {
	fmt.Println(Name, "version", Version)
}

func OnStart(g *api.API) error {
	fmt.Println("UPlay OnStart called")

	return nil
}
