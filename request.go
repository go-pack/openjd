package openjd

type Request interface {
	GetApiMethodName() string
	GetApiParas() string
}

type Params map[string]interface{}
func NewParams() Params {
	p := make(Params)
	return p
}
func (p Params) Set(key string, value interface{}) {
	p[key] = value
}
func (p Params) GetParams() ( map[string]interface{}) {
	return p
}
