package plugin

import (
	"errors"
	"log"
	"plugin"
	"sync"
)

type PluginStatus byte

const (
	PluginDisabled    PluginStatus = 0
	PluginEnabled     PluginStatus = 1
	PluginStarting    PluginStatus = 2
	PluginInitialized PluginStatus = 3
	PluginError       PluginStatus = 4
)

func (s PluginStatus) Name() string {
	switch s {
	case 0:
		return "disabled"
	case 1:
		return "starting"
	case 2:
		return "enabled"
	case 3:
		return "initialized"
	case 4:
		return "error"
	}

	return "invalid"
}

type Plugin struct {
	Status PluginStatus
	IsGone bool

	ID               string
	Name             string
	Version          string
	Authors          []string
	ShortDescription string
	LongDescription  string

	OnStart func() error
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
		plugin.OnStart = sym.(func() error)
	} else {
		return nil, errors.New("Plugin does not have a Name field")
	}

	return &plugin, nil
}

func (p Plugin) Start(wg *sync.WaitGroup) error {
	defer wg.Done()

	var ret error

	defer func() {
		// If the plugin panics, we should return an error, rather than crashing everything
		if err := recover(); err != nil {
			p.Status = PluginError
			_, ok := err.(error)
			if !ok {
				panic("Returned type is not an error!")
			}

			ret = err.(error)
		}
	}()

	p.Status = PluginStarting

	// Call the OnStart() function in the plugin
	ret = p.OnStart()

	if ret != nil {
		// If an error was returned, set the status
		p.Status = PluginError
	} else {
		// Otherwise, it enabled successfully
		p.Status = PluginInitialized
	}

	return ret
}
