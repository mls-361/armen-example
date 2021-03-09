/*
------------------------------------------------------------------------------------------------------------------------
####### job ####### (c) 2020-2021 mls-361 ########################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package job

import (
	"github.com/mls-361/armen-sdk/jw"
	"github.com/mls-361/logger"

	"github.com/mls-361/armen-example/internal/resources"
)

type (
	// Job AFAIRE.
	Job struct {
		*jw.Job
		Logger    logger.Logger
		Resources *resources.Resources
	}
)

// New AFAIRE.
func New(job *jw.Job, logger logger.Logger, resources *resources.Resources) *Job {
	return &Job{
		Job:       job,
		Logger:    logger,
		Resources: resources,
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
