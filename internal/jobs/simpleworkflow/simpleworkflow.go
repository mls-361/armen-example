/*
------------------------------------------------------------------------------------------------------------------------
####### simpleworkflow ####### (c) 2020-2021 mls-361 ############################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package simpleworkflow

import (
	"github.com/mls-361/armen-sdk/jw"
	"github.com/mls-361/datamap"
	"github.com/mls-361/uuid"

	"github.com/mls-361/armen-example/internal/job"
	"github.com/mls-361/armen-example/internal/jobs"
)

type (
	// Job AFAIRE.
	Job struct {
		*job.Job
	}
)

// New AFAIRE.
func New(job *job.Job) *Job {
	return &Job{Job: job}
}

// Run AFAIRE.
func (j *Job) Run() *jw.Result {
	const (
		firstStep  = "one"
		secondStep = "two"
		thirdStep  = "three"
	)

	namespace := j.Resources.Plugin.Name()
	succeeded := jw.StatusSucceeded.String()

	allSteps := map[string]*jw.Step{
		firstStep: jw.NewStep(
			namespace,
			jobs.One,
			nil,
			datamap.DataMap{succeeded: secondStep},
		),
		secondStep: jw.NewStep(
			namespace,
			jobs.Two,
			nil,
			datamap.DataMap{succeeded: datamap.DataMap{"step": thirdStep}},
		),
		thirdStep: jw.NewStep(
			namespace,
			jobs.Three,
			nil,
			nil,
		),
	}

	wf := jw.NewWorkflow(
		uuid.New(),
		jobs.SimpleWorkflow,
		jobs.SimpleWorkflow,
		jw.PriorityLow,
		firstStep,
		allSteps,
	)

	if err := j.Resources.Model().InsertWorkflow(wf); err != nil {
		return jw.Failed(err)
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/