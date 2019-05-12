package gridlock

import (
	"log"
	"net/http"
	"encoding/json"
	"strings"

	"github.com/c2h5oh/datasize"
	"github.com/gorilla/mux"

	"github.com/Stumblinbear/gridlock/api"
	"gridlock/plugin"
)

var version string

type Gridlock struct {
	router *mux.Router

	pm plugin.Manager
	api api.API // Should I create an API struct for each plugin?

	Store Store

	// Sends notifications out to anyone connected to the websocket
	notify chan api.Notification
}

type Store struct {
	// The current machine is always named "self"
	Hosts map[string]*api.Host `json:"hosts"`

	Metadata map[string]*api.Metadata `json:"metadata"`
}

func NewGridlock() *Gridlock {
	log.Printf("Starting Gridlock v%s...\n", version)

	log.Println("Fetching system information...")

	system, err := GetSystemInfo()
	if err != nil {
		panic(err)
	}

	log.Println("")

	log.Printf("Unique ID: %s\n", system.ID)

	log.Printf("Hostname: %s (%s, %s %s)\n",
		system.Hostname,
		system.Platform.ID,
		system.Platform.Name,
		system.Platform.Version)

	log.Printf("CPU: %s\n", system.CPU.Name)

	log.Printf("RAM Total: %s, Free: %s, Used: %s\n",
		datasize.ByteSize(system.RAM.Total).HR(),
		datasize.ByteSize(system.RAM.Free).HR(),
		datasize.ByteSize(system.RAM.Used).HR())

	log.Println("")

	self := api.Host{
		Remote: false,
		System: system,

		Platforms: map[string]api.Platform{},
		Launchers: map[string]api.Launcher{},
	}

	g := Gridlock{
		router: mux.NewRouter().StrictSlash(true),

		pm: plugin.NewManager(),

		Store: Store{
			Hosts: map[string]*api.Host{
				"self": &self,
			},
			Metadata: map[string]*api.Metadata{},
		},
	}

	g.api = api.API{
		AddLauncher: func(l api.Launcher) {
			self.Launchers[IDify(l.Name)] = l
		},
		AddPlatform: func(p api.Platform) {
			self.Platforms[IDify(p.Name)] = p
		},
	}

	g.api.AddLauncher(api.Launcher{
		Name: "Direct",
		Description: "Launch the game on this machine.",

		CanStart: func(gq api.GameQuery) bool {
			return true
		},
		StartGame: func(gq api.GameQuery) error {
			// Poke the platform to start the game directly
			return nil
		},
	})

	return &g
}

func IDify(name string) string {
	name = strings.ToLower(name)

	name = strings.Replace(name, " ", "-", -1)

	return name
}

type Callback func(*http.Request) (int, interface{})

func (g Gridlock) AddEndpoint(path string, cb Callback) {
	log.Println("Registered endpoint:", path)

	g.router.HandleFunc("/" + path, func(w http.ResponseWriter, r *http.Request) {
		status, payload := cb(r)

		switch payload.(type) {
			case error:
				err := payload.(error)

				if status == 0 {
					status = 500				
				}
					
				payload = map[string]string {
					"message": err.Error(),
				}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)

		if payload != nil {
			payload, err := json.MarshalIndent(payload, "", "  ")

			if err != nil {
				panic(err)
			}
	
			w.Write(payload)
		}
	})
}

func (g Gridlock) RunForever() {
	// Start the HTTP listener as soon as possible.
	RegisterRoutes(&g)

	go func() {
		log.Fatal(http.ListenAndServe(*addr, g.router))
	}()

	// After it's up, initialize plugins.
	err := g.pm.RefreshPlugins()
	if err != nil {
		panic(err)
	}

	// Initialize the plugin manager
	go g.pm.Initialize(&g.api)

	select{}
}