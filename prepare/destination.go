package prepare

type DestinationReq struct {
	doName string
	uri    string
	body   map[string]string
	method string
	param  map[string]string
}

func (d *DestinationReq) Param() map[string]string {
	return d.param
}

func (d *DestinationReq) SetParam(param map[string]string) {
	d.param = param
}

func (d *DestinationReq) Method() string {
	return d.method
}

func (d *DestinationReq) SetMethod(method string) {
	d.method = method
}

func (d *DestinationReq) Body() map[string]string {
	return d.body
}

func (d *DestinationReq) SetBody(body map[string]string) {
	d.body = body
}

func (d *DestinationReq) Uri() string {
	return d.uri
}

func (d *DestinationReq) SetUri(uri string) {
	d.uri = uri
}

func (d *DestinationReq) DoName() string {
	return d.doName
}

func (d *DestinationReq) SetDoName(doName string) {
	d.doName = doName
}

func NewDestinationReq(doName string, uri string, body map[string]string, method string, param map[string]string) *DestinationReq {
	return &DestinationReq{doName: doName, uri: uri, body: body, method: method, param: param}
}
