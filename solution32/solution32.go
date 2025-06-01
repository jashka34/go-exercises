package solution32

import (
	"errors"
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
	// fmt.Println("--------\njob: ", job)
	if job == nil || job.Dicts == nil {
		// fmt.Println("ret1")
		// fmt.Println("ret11")
		return job, errNotEnoughDicts
	}
	if len(job.Dicts) < 2 {
		// fmt.Println("ret2")
		return job, errNotEnoughDicts
	}
	job.Merged = make(map[string]string)
	// defer finish(&job)
	for _, v := range job.Dicts {
		if v != nil {
			// fmt.Println(i, ":", v)
			// maps.Copy(job.Merged, v)
			for ii, vv := range v {
				// fmt.Println("  ", ii, ":", vv)
				job.Merged[ii] = vv
			}
		} else {
			// fmt.Println("ret3")
			return job, errNilDict
		}
	}
	return job, nil
}
