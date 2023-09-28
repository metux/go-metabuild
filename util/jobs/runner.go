package jobs

import (
	"fmt"
	"log"
	// "sync"
)

type Runner struct {
	Jobs map[JobId]*RunnerJob
}

func NewRunner() Runner {
	r := Runner{}
	r.Jobs = make(map[JobId]*RunnerJob)
	return r
}

func (r Runner) AddJob(id JobId, job Job) error {
	j, err := NewJob(id, job)
	if err != nil {
		log.Println("AddJob: error creating job runner: ", err)
		return err
	}

	r.Jobs[id] = j

	for _, walk := range j.Sub {
		r.AddJob(walk.JobId(), walk)
	}

	return err
}

// resolve dependencies
func (r Runner) Resolve() error {
	for id, walk := range r.Jobs {
		for _, dwalk := range walk.Job.JobDepends() {
			if d1, ok := r.Jobs[dwalk]; ok {
				walk.Depends = append(walk.Depends, d1)
			} else {
				err := fmt.Errorf("missing dep %s for %s", dwalk, id)
				return err
			}
		}
		for _, subWalk := range walk.Sub {
			subId := subWalk.JobId()
			if d1, ok := r.Jobs[subId]; ok {
				walk.Depends = append(walk.Depends, d1)
			} else {
				err := fmt.Errorf("missing sub dep %s for %s", subId, id)
				return err
			}
		}
	}
	return nil
}

func (r Runner) Scan() (RunnerJobList, RunnerJobList, RunnerJobList) {
	doneList := RunnerJobList{}
	runnableList := RunnerJobList{}
	waitList := RunnerJobList{}

	for _, walk := range r.Jobs {
		if walk.Done {
			doneList = append(doneList, walk)
		} else if walk.Runnable() {
			runnableList = append(runnableList, walk)
		} else {
			waitList = append(waitList, walk)
		}
	}
	return doneList, runnableList, waitList
}

// FIXME: parallel runs
func (r Runner) Run() error {
	if err := r.Resolve(); err != nil {
		return err
	}

	for true {
		done, runnable, waiting := r.Scan()
		log.Printf("done: %d runnable: %d waiting: %d\n", len(done), len(runnable), len(waiting))

		if len(runnable) == 0 {
			log.Println("no more waiting. finished")
			if len(waiting) != 0 {
				return fmt.Errorf("some jobs cant run: %d", len(waiting))
			}
			return nil
		}

		for _, b := range runnable {
			log.Println("Running job", b.Id)
			if err := b.Job.JobRun(); err != nil {
				return err
			}
			b.Done = true
			log.Println("Done job", b.Id)
		}
	}
	return nil
}
