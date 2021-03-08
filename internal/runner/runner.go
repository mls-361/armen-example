/*
------------------------------------------------------------------------------------------------------------------------
####### runner ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package runner

import (
	"github.com/mls-361/armen-sdk/jw"
	"github.com/mls-361/armen-sdk/logger"
	"github.com/mls-361/failure"
	"github.com/mls-361/minikit"

	_job "github.com/mls-361/armen-example/internal/job"
	"github.com/mls-361/armen-example/internal/jobs"
	"github.com/mls-361/armen-example/internal/jobs/donothing"
	"github.com/mls-361/armen-example/internal/jobs/donothingwithkey"
	"github.com/mls-361/armen-example/internal/jobs/one"
	"github.com/mls-361/armen-example/internal/jobs/simpleworkflow"
	"github.com/mls-361/armen-example/internal/jobs/three"
	"github.com/mls-361/armen-example/internal/jobs/two"
	"github.com/mls-361/armen-example/internal/resources"
)

type (
	// Runner AFAIRE.
	Runner struct {
		*minikit.Base
		resources *resources.Resources
	}
)

// New AFAIRE.
func New(pluginName string, resources *resources.Resources) *Runner {
	return &Runner{
		Base:      minikit.NewBase(pluginName+".runner", ""),
		resources: resources,
	}
}

// RunJob AFAIRE.
func (cr *Runner) RunJob(job *jw.Job, logger logger.Logger) *jw.Result {
	j := _job.New(job, logger, cr.resources)

	switch j.Type {
	case jobs.DoNothing:
		return donothing.New(j).Run()
	case jobs.DoNothingWithKey:
		return donothingwithkey.New(j).Run()
	case jobs.One:
		return one.New(j).Run()
	case jobs.SimpleWorkflow:
		return simpleworkflow.New(j).Run()
	case jobs.Three:
		return three.New(j).Run()
	case jobs.Two:
		return two.New(j).Run()
	default:
		return jw.Failed(
			failure.New(nil).
				Set("namespace", j.Namespace).
				Set("type", j.Type).
				Msg("this type of job does not exist for this namespace"), /////////////////////////////////////////////
		)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
