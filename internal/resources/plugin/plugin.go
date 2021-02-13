/*
------------------------------------------------------------------------------------------------------------------------
####### plugin ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package plugin

import (
	"github.com/mls-361/armen-sdk/plugin"
	"github.com/mls-361/minikit"

	"github.com/mls-361/armen-example/internal/resources"
)

type (
	// Plugin AFAIRE.
	Plugin struct {
		*plugin.Plugin
		resources *resources.Resources
	}
)

// New AFAIRE.
func New(name, version, builtAt string, resources *resources.Resources) *Plugin {
	plugin := &Plugin{
		Plugin:    plugin.New(name, version, builtAt, resources.Components),
		resources: resources,
	}

	resources.Plugin = plugin

	return plugin
}

// Build AFAIRE.
func (cp *Plugin) Build(m *minikit.Manager) error {
	if err := cp.Plugin.Build(m); err != nil {
		return err
	}

	// On profite de la construction du composant 'Plugin'
	// pour créer le composant 'Logger'.
	// Cela est possible, car la construction du composant 'Plugin' dépend
	// de celle du composant 'Logger' de l'application.
	cp.resources.Logger = cp.resources.Components.Logger().CreateLogger(cp.ID(), cp.Name())

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
