package jobs

type RunnerJobRef *RunnerJob

type RunnerJob struct {
	Id   string
	Job  Job
	Done bool
	// FIXME: lock/mutex
	Depends []RunnerJobRef
	Sub     []Job
}

type RunnerJobList []RunnerJobRef

func NewJob(id JobId, job Job) (RunnerJobRef, error) {
	if err := job.JobPrepare(id); err != nil {
		return nil, err
	}

	sub, err := job.JobSub()
	if err != nil {
		return nil, err
	}

	rj := RunnerJob{
		Id:  id,
		Job: job,
		Sub: sub,
	}
	return &rj, nil
}

func (rj RunnerJob) Run() error {
	return rj.Job.JobRun()
}

func (rj RunnerJob) Runnable() bool {
	for _, walk := range rj.Depends {
		if !walk.Done {
			return false
		}
	}
	return true
}
