/*
------------------------------------------------------------------------------------------------------------------------
####### main ####### (c) 2020-2021 mls-361 ######################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"github.com/mls-361/armen-sdk/components"
	"github.com/mls-361/minikit"

	"github.com/mls-361/armen-example/internal/msghandler"
	"github.com/mls-361/armen-example/internal/resources"
	"github.com/mls-361/armen-example/internal/resources/factory"
	"github.com/mls-361/armen-example/internal/resources/plugin"
	"github.com/mls-361/armen-example/internal/runner"
)

const (
	_pluginName = "example"
)

var (
	_version string
	_builtAt string
)

// Export AFAIRE.
func Export(m *minikit.Manager, cs components.Components) error {
	rs := resources.New(cs)

	return m.AddComponents(
		factory.New(_pluginName, rs),
		msghandler.New(_pluginName, rs),
		plugin.New(_pluginName, _version, _builtAt, rs),
		runner.New(_pluginName, rs),
	)
}

func main() {
	_ = Export(nil, nil) // avoid linter errors ////////////////////////////////////////////////////////////////////////
}

/*
######################################################################################################## @(°_°)@ #######
*/
