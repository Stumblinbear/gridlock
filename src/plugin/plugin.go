package plugin

import (
	"errors"
	"log"
	"plugin"
	"sync"

	"github.com/Stumblinbear/gridlock/api"
)

type Status byte

const (
	Disabled Status = iota
	Enabled
	Starting
	Initialized
	Error
)

var statuses = map[Status]string{
	Disabled:    "DISABLED",
	Enabled:     "ENABLED",
	Starting:    "STARTING",
	Initialized: "INITIALIZED",
	Error:       "ERROR",
}

// String prints the string version of the Op consts
func (e Status) String() string {
	if status, found := statuses[e]; found {
		return status
	}
	return "???"
}

type Plugin struct {
	Status Status
	IsGone bool

	ID               string
	Name             string
	Version          string
	Authors          []string
	ShortDescription string
	LongDescription  string

	OnStart func(*api.API) error
}

func LoadPlugin(id string, file string) (*Plugin, error) {
	plug, err := plugin.Open(file)
	if err != nil {
		return nil, err
	}

	plugin := Plugin{
		ID: id,
	}

	if sym, err := plug.Lookup("Name"); err == nil {
		plugin.Name = *sym.(*string)
	} else {
		return nil, errors.New("Plugin does not have a Name field")
	}

	if sym, err := plug.Lookup("Version"); err == nil {
		plugin.Version = *sym.(*string)
	} else {
		return nil, errors.New("Plugin does not have a Version field")
	}

	if sym, err := plug.Lookup("Authors"); err == nil {
		plugin.Authors = *sym.(*[]string)
	} else if sym, err := plug.Lookup("Author"); err == nil {
		plugin.Authors = []string{*sym.(*string)}
	}

	if sym, err := plug.Lookup("Description"); err == nil {
		plugin.ShortDescription = *sym.(*string)
		plugin.LongDescription = *sym.(*string)
	}

	if sym, err := plug.Lookup("ShortDescription"); err == nil {
		plugin.ShortDescription = *sym.(*string)
	}

	if sym, err := plug.Lookup("LongDescription"); err == nil {
		plugin.LongDescription = *sym.(*string)
	}

	if plugin.ShortDescription == "" && plugin.LongDescription != "" {
		plugin.ShortDescription = plugin.LongDescription[:50]
	}

	if plugin.ShortDescription != "" && plugin.LongDescription == "" {
		plugin.LongDescription = plugin.ShortDescription
	}

	if len(plugin.ShortDescription) > 50 {
		log.Println("Short descriptin is too long. Truncating.")
		plugin.ShortDescription = plugin.ShortDescription[:50]
	}

	if sym, err := plug.Lookup("OnStart"); err == nil {
		plugin.OnStart = sym.(func(*api.API) error)
	} else {
		return nil, errors.New("Plugin does not have a Name field")
	}

	return &plugin, nil
}

func (p Plugin) Start(api *api.API, wg *sync.WaitGroup) error {
	defer wg.Done()

	var ret error

	defer func() {
		// If the plugin panics, we should return an error, rather than crashing everything
		if err := recover(); err != nil {
			p.Status = Error
			_, ok := err.(error)
			if !ok {
				panic("Returned type is not an error!")
			}

			ret = err.(error)
		}
	}()

	p.Status = Starting

	// Call the OnStart() function in the plugin
	ret = p.OnStart(api)

	if ret != nil {
		// If an error was returned, set the status
		p.Status = Error
	} else {
		// Otherwise, it enabled successfully
		p.Status = Initialized
	}

	return ret
}
