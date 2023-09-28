package jobs

type BaseJob struct {
	Id JobId
}

func (j BaseJob) JobPrepare(id JobId) error {
	return nil
}

func (j BaseJob) JobRun() error {
	return nil
}

func (j BaseJob) JobId() string {
	return string(j.Id)
}

func (j BaseJob) JobDepends() []JobId {
	return []JobId{}
}

func (j BaseJob) JobSub() ([]Job, error) {
	return []Job{}, nil
}

func MakeBaseJob(id JobId) BaseJob {
	return BaseJob{id}
}
