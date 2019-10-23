package prepare

/**
doc详情
*/
type Document struct {
	company Company
	job     Job
	time    string
}

func NewDocument(company Company, job Job, time string) *Document {
	return &Document{company: company, job: job, time: time}
}

type Company struct {
	name     string
	link     string
	desc     string
	location string
	types    string
}

type Job struct {
	name   string
	salary string
	desc   string
}
