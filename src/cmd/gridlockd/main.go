package main

import (
	"log"
	"os"

	"gridlock"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.Fatal("An error occured: ", err)
		}
	}()

	log.SetOutput(os.Stdout)

	g := gridlock.NewGridlock()

	g.RunForever()
}
