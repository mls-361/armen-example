/*
------------------------------------------------------------------------------------------------------------------------
####### factory ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package factory

import (
	"github.com/mls-361/armen-sdk/jw"
	"github.com/mls-361/armen-sdk/message"
	"github.com/mls-361/failure"
	"github.com/mls-361/minikit"

	"github.com/mls-361/armen-example/internal/jobs"
	"github.com/mls-361/armen-example/internal/resources"
)

type (
	// Factory AFAIRE.
	Factory struct {
		*minikit.Base
		namespace string
		resources *resources.Resources
	}
)

// New AFAIRE.
func New(pluginName string, resources *resources.Resources) *Factory {
	cf := &Factory{
		Base:      minikit.NewBase(pluginName+".factory", ""),
		namespace: pluginName,
		resources: resources,
	}

	resources.Factory = cf

	return cf
}

func (cf *Factory) createJob(id, name, _type, origin string, priority jw.Priority, key *string) *jw.Job {
	return jw.NewJob(id, name, cf.namespace, _type, origin, priority, key)
}

func (cf *Factory) createJobDoNothing(msg *message.Message) {
	job := cf.createJob(
		msg.ID,
		jobs.DoNothing,
		jobs.DoNothing,
		msg.Publisher,
		jw.NoPriority,
		nil,
	)

	_ = cf.resources.Model().InsertJob(job)
}

func (cf *Factory) createJobDoNothingWithKey(msg *message.Message) {
	key := "JustAKey"

	job := cf.createJob(
		msg.ID,
		jobs.DoNothingWithKey,
		jobs.DoNothingWithKey,
		msg.Publisher,
		jw.NoPriority,
		&key,
	)

	_ = cf.resources.Model().InsertJob(job)
}

func (cf *Factory) createJobSimpleWorkflow(msg *message.Message) {
	job := cf.createJob(
		msg.ID,
		jobs.SimpleWorkflow,
		jobs.SimpleWorkflow,
		msg.Publisher,
		jw.NoPriority,
		nil,
	)

	_ = cf.resources.Model().InsertJob(job)
}

// CreateJob AFAIRE.
func (cf *Factory) CreateJob(jt string, msg *message.Message) error {
	switch jt {
	case jobs.DoNothing:
		cf.createJobDoNothing(msg)
	case jobs.DoNothingWithKey:
		cf.createJobDoNothingWithKey(msg)
	case jobs.SimpleWorkflow:
		cf.createJobSimpleWorkflow(msg)
	default:
		return failure.New(nil).
			Set("type", jt).
			Set("namespace", cf.namespace).
			Msg("this type of job is not known in this namespace") /////////////////////////////////////////////////////
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
