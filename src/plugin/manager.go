package plugin

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"gridlock/api"
)

type PluginManager struct {
	initialized bool

	plugins map[string]Plugin
}

func (pm PluginManager) Initialize() {
	if pm.initialized {
		panic("Plugin system already started!")
	}

	var wg sync.WaitGroup

	for _, p := range pm.plugins {
		pm.InitPlugin(&wg, p.ID)
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
func (pm PluginManager) RefreshPlugins(api *api.API) error {
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

func (pm PluginManager) HasPlugin(id string) bool {
	_, ok := pm.plugins[id]

	return ok
}

func (pm PluginManager) EnablePlugin(wg *sync.WaitGroup, id string) {
	if !pm.HasPlugin(id) {
		panic("Unknown plugin ID!")
	}

	plugin := pm.plugins[id]

	if plugin.Status != PluginDisabled {
		panic("Plugin is already enabled!")
	}

	plugin.Status = PluginEnabled

	// If the plugin manager has already been initialized, this plugin should be immediately started
	if pm.initialized {
		pm.InitPlugin(wg, id)
	}
}

func (pm PluginManager) InitPlugin(wg *sync.WaitGroup, id string) {
	if !pm.HasPlugin(id) {
		panic("Unknown plugin ID!")
	}

	plugin := pm.plugins[id]

	if plugin.Status != PluginDisabled {
		panic("Plugin can only be started when in the disabled state!")
	}

	wg.Add(1)
	go plugin.Start(wg)
}

func NewPluginManager() *PluginManager {
	manager := PluginManager{
		initialized: false,
		plugins:     make(map[string]Plugin),
	}

	return &manager
}
