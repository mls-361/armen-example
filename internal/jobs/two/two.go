/*
------------------------------------------------------------------------------------------------------------------------
####### two ####### (c) 2020-2021 mls-361 ########################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package two

import (
	"github.com/mls-361/armen-sdk/jw"

	"github.com/mls-361/armen-example/internal/job"
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
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
