package requests

import (
	"encoding/json"
	"github.com/go-pack/openjd"
)

type JdUnionOpenOrderRequest struct {
	openjd.Params
}

func NewJdUnionOpenOrderRequest() *JdUnionOpenOrderRequest {
	return &JdUnionOpenOrderRequest{Params: openjd.NewParams()}
}

func (request *JdUnionOpenOrderRequest) GetApiMethodName() string {
	return "jd.union.open.order.query"
}

func (request *JdUnionOpenOrderRequest) GetApiParas() string {
	mp := make(map[string]interface{})
	mp["orderReq"] = request.GetParams()
	marshal, _ := json.Marshal(mp)
	return string(marshal)
}



func (request *JdUnionOpenOrderRequest) SetTime(StartTime string) {
	request.Set("time", StartTime)
}

func (request *JdUnionOpenOrderRequest) SetPageNo(PageNo int) {
	request.Set("pageNo", PageNo)
}
func (request *JdUnionOpenOrderRequest) SetChildUnionId(childUnionId string) {
	request.Set("childUnionId", childUnionId)
}

func (request *JdUnionOpenOrderRequest) SetPageSize(PageSize int) {
	request.Set("pageSize", PageSize)
}

func (request *JdUnionOpenOrderRequest) SetType(queryType int) {
	request.Set("type", queryType)
}

func (request *JdUnionOpenOrderRequest) SetKey(key string) {
	request.Set("key", key)
}
