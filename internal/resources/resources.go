/*
------------------------------------------------------------------------------------------------------------------------
####### resources ####### (c) 2020-2021 mls-361 #################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package resources

import (
	"github.com/mls-361/armen-sdk/components"
	"github.com/mls-361/armen-sdk/message"
	"github.com/mls-361/logger"
)

type (
	// Factory AFAIRE.
	Factory interface {
		CreateJob(jt string, msg *message.Message) error
	}

	// Resources AFAIRE.
	Resources struct {
		components.Components
		Factory Factory
		Logger  logger.Logger
		Plugin  components.Plugin
	}
)

// New AFAIRE.
func New(components components.Components) *Resources {
	return &Resources{
		Components: components,
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
