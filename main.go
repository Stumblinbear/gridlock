package main

import (
	"log"
	"os"

	"gridlock/api"
	"gridlock/plugin"
	"gridlock/sys"

	"github.com/c2h5oh/datasize"
)

var version string

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.Fatal("An error occured: ", err)
		}
	}()

	log.SetOutput(os.Stdout)

	log.Printf("Starting Gridlock v%s...\n", version)

	log.Println("Fetching system information...")

	info, err := sys.GetSystemInfo()
	if err != nil {
		panic(err)
	}

	log.Println("")

	log.Printf("Unique ID: %s\n", info.ID)

	log.Printf("Hostname: %s (%s, %s %s)\n",
		info.Hostname,
		info.Platform.ID,
		info.Platform.Name,
		info.Platform.Version)

	log.Printf("CPU: %s\n", info.CPU.Name)

	log.Printf("RAM Total: %s, Free: %s, Used: %s\n",
		datasize.ByteSize(info.RAM.Total).HR(),
		datasize.ByteSize(info.RAM.Free).HR(),
		datasize.ByteSize(info.RAM.Used).HR())

	log.Println("")

	pm := plugin.NewPluginManager()

	api := api.Create()
	
	err = pm.RefreshPlugins(&api)
	if err != nil {
		panic(err)
	}

	pm.Initialize()
}
