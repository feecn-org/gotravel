package prepare

/**
招聘页详情
*/
type Document struct {
	company string
	jobName string
}

func NewDocument(company string, jobName string) *Document {
	return &Document{company: company, jobName: jobName}
}

type company struct {
}

type job struct {
}
