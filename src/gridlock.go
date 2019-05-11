package main

import (
	"log"

	"github.com/Stumblinbear/gridlock/api"
	"gridlock/sys"
	"gridlock/plugin"
)

type Gridlock struct {
	system sys.Info

	pm plugin.Manager
	api api.API // Should I create an API struct for each plugin?
}

func NewGridlock(system sys.Info) *Gridlock {
	g := Gridlock{
		system: system,
		pm: plugin.NewManager(),
		api: api.Create(),
	}

	return &g
}

func (g Gridlock) RunForever() {
	err := g.pm.RefreshPlugins()
	if err != nil {
		panic(err)
	}

	g.pm.Initialize(&g.api)

	log.Println("Listening for commands...")

	select{}
}