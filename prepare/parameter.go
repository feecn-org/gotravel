package prepare

/**
zhaopin 参数: first page
https://fe-api.zhaopin.com/c/i/sou?
pageSize=90&
cityId=530&
workExperience=-1&
education=-1&
companyType=-1&
employmentType=-1&
jobWelfareTag=-1&
kw=java&
kt=3&
_v=0.45623841&
x-zp-page-request-id=7ebedd4821b54a5a8f0b8471356e5dfb-1571280451080-760373&
x-zp-client-id=f074fd10-6474-450f-8f4c-3826a9114f39

zhaopin 参数: second page
https://fe-api.zhaopin.com/c/i/sou?
start=90&
pageSize=90&
cityId=530&
workExperience=-1&
education=-1&
companyType=-1&
employmentType=-1&
jobWelfareTag=-1&
kw=java&
kt=3&
_v=0.45623841&
x-zp-page-request-id=7ebedd4821b54a5a8f0b8471356e5dfb-1571280451080-760373&
x-zp-client-id=f074fd10-6474-450f-8f4c-3826a9114f39
*/

type Parameter struct {
	getMap  map[string]string
	postMap map[string]string
}

func (p *Parameter) PostMap() map[string]string {
	return p.postMap
}

func (p *Parameter) SetPostMap(postMap map[string]string) {
	p.postMap = postMap
}

func (p *Parameter) GetMap() map[string]string {
	return p.getMap
}

func (p *Parameter) SetGetMap(getMap map[string]string) {
	p.getMap = getMap
}

func NewParameter(getMap map[string]string, postMap map[string]string) *Parameter {
	return &Parameter{getMap: getMap, postMap: postMap}
}
