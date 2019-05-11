package plugin

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/Stumblinbear/gridlock/api"
)

type Manager struct {
	initialized bool

	plugins map[string]Plugin
}

func NewManager() Manager {
	return Manager{
		initialized: false,
		plugins:     make(map[string]Plugin),
	}
}

func (pm Manager) Initialize(gapi *api.API) {
	if pm.initialized {
		panic("Plugin system already started!")
	}

	var wg sync.WaitGroup

	for _, p := range pm.plugins {
		pm.InitPlugin(gapi, p.ID, &wg)
	}

	log.Println("Waiting for plugins to initialize...")
	wg.Wait()
	log.Println("Done")

	pm.initialized = true
}

/**
This will search for changes in the plugins folder
and apply them to the plugin manager.
*/
func (pm Manager) RefreshPlugins() error {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}

	pluginFolder := filepath.Join(dir, "plugins") + "/"

	log.Println("Refreshing plugins...")

	files, err := ioutil.ReadDir(pluginFolder)
	if err != nil {
		return err
	}

	for _, f := range files {
		file := f.Name()

		// We should only load plugin files
		if filepath.Ext(file) != ".so" {
			continue
		}

		// Trim the extension
		id := strings.TrimSuffix(file, ".so")

		// If the plugin is already detected, ignore it
		if pm.HasPlugin(id) {
			continue
		}

		plugin, err := LoadPlugin(id, pluginFolder+file)

		if err != nil {
			log.Println("Failed to load plugin:", id)
			log.Println(" ", err)
			continue
		}

		log.Println("Loaded plugin:", plugin.Name, "version", plugin.Version, "by", plugin.Authors, "(", plugin.ShortDescription, ")")

		pm.plugins[id] = *plugin
	}

	return nil
}

func (pm Manager) HasPlugin(id string) bool {
	_, ok := pm.plugins[id]

	return ok
}

func (pm Manager) EnablePlugin(gapi *api.API, id string, wg *sync.WaitGroup) {
	if !pm.HasPlugin(id) {
		panic("Unknown plugin ID!")
	}

	plugin := pm.plugins[id]

	if plugin.Status != Disabled {
		panic("Plugin is already enabled!")
	}

	plugin.Status = Enabled

	// If the plugin manager has already been initialized, this plugin should be immediately started
	if pm.initialized {
		pm.InitPlugin(gapi, id, wg)
	}
}

func (pm Manager) InitPlugin(gapi *api.API, id string, wg *sync.WaitGroup) {
	if !pm.HasPlugin(id) {
		panic("Unknown plugin ID!")
	}

	plugin := pm.plugins[id]

	if plugin.Status != Disabled {
		panic("Plugin can only be started when in the disabled state!")
	}

	wg.Add(1)
	go plugin.Start(gapi, wg)
}
