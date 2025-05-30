package solution32

import (
	"errors"
	"fmt"
)

// MergeDictsJob is a job to merge dictionaries into a single dictionary.
type MergeDictsJob struct {
	Dicts      []map[string]string
	Merged     map[string]string
	IsFinished bool
}

// errors
var (
	errNotEnoughDicts = errors.New("at least 2 dictionaries are required")
	errNilDict        = errors.New("nil dictionary")
)

func ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error) {
	defer job.IsFinished = true
	fmt.Println("job: ", job)
	if job == nil || job.Dicts == nil {
		// fmt.Println("ret1")
		// fmt.Println("ret11")
		return job, errNotEnoughDicts
	}
	if len(job.Dicts) < 2 {
		fmt.Println("ret2")
		return job, errNotEnoughDicts
	}
	var ret MergeDictsJob
	for i, v := range job.Dicts {
		fmt.Println(i, ":", v)
	}
	return &ret, nil
}
