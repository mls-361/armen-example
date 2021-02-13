/*
------------------------------------------------------------------------------------------------------------------------
####### msghandler ####### (c) 2020-2021 mls-361 ################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package msghandler

import (
	"fmt"
	"strings"

	"github.com/mls-361/armen-sdk/message"
	"github.com/mls-361/minikit"

	"github.com/mls-361/armen-example/internal/resources"
)

type (
	// MsgHandler AFAIRE.
	MsgHandler struct {
		*minikit.Base
		pluginName string
		cjPrefix   string
		resources  *resources.Resources
	}
)

// New AFAIRE.
func New(pluginName string, resources *resources.Resources) *MsgHandler {
	return &MsgHandler{
		Base:       minikit.NewBase("msghandler."+pluginName, "msghandler."+pluginName),
		pluginName: pluginName,
		cjPrefix:   fmt.Sprintf("create.job.%s.", pluginName),
		resources:  resources,
	}
}

// Dependencies AFAIRE.
func (cmh *MsgHandler) Dependencies() []string {
	return []string{
		"bus",
	}
}

// Build AFAIRE.
func (cmh *MsgHandler) Build(_ *minikit.Manager) error {
	return cmh.resources.Bus().Subscribe(
		cmh.consume,
		`create[.]job[.]`+cmh.pluginName+`[.].+`,
	)
}

func (cmh *MsgHandler) createJob(msg *message.Message) {
	if err := cmh.resources.Factory.CreateJob(strings.TrimPrefix(msg.Topic, cmh.cjPrefix), msg); err != nil {
		cmh.resources.Logger.Error( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Error when creating a job",
			"message", msg.ID,
			"publisher", msg.Publisher,
			"reason", err.Error(),
		)
	}
}

func (cmh *MsgHandler) consume(msg *message.Message) {
	cmh.resources.Logger.Debug( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		"Consume message",
		"id", msg.ID,
		"topic", msg.Topic,
		"publisher", msg.Publisher,
	)

	if strings.HasPrefix(msg.Topic, cmh.cjPrefix) {
		cmh.createJob(msg)
		return
	}

	switch msg.Topic {
	default:
		cmh.resources.Logger.Error( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"The topic of this message is not valid",
			"id", msg.ID,
			"topic", msg.Topic,
			"publisher", msg.Publisher,
		)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
