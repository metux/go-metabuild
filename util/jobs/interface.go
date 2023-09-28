package jobs

type JobId = string

type Job interface {
	// called on adding a job to a runner (not parallelized)
	JobPrepare(id JobId) error

	// run the job (potentially parallelized)
	JobRun() error

	JobId() JobId
	JobDepends() []JobId
	JobSub() ([]Job, error)
}

type JobMap = map[JobId]Job
