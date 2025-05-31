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

func finish(job *MergeDictsJob) {
	job.IsFinished = true
}

func ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error) {
	defer finish(job)
	fmt.Println("--------\njob: ", job)
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
	ret.Merged = make(map[string]string)
	defer finish(&ret)
	for i, v := range job.Dicts {
		if v != nil {
			fmt.Println(i, ":", v)
			for ii, vv := range v {
				fmt.Println("  ", ii, ":", vv)
				ret.Merged[ii] = vv
				fmt.Println("  merged!")
			}
		} else {
			fmt.Println("ret3")
			return &ret, errNilDict
		}
	}
	return &ret, nil
}
