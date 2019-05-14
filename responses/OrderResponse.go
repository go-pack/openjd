package responses

import (
	"encoding/json"
)

type OrderResponse struct {
	BaseResponse
	FinishTime uint64 `json:"finishTime"`
	OrderId uint64 `json:"orderId"`
	OrderTime uint64 `json:"orderTime"`
	ParentId uint64 `json:"parentId"`
	//实际计算佣金的金额。订单完成后，会将误扣除的运费券金额更正。如订单完成后发生退款，此金额会更新。
	ActualCosPrice float32 `json:"actualCosPrice"`
	//推客获得的实际佣金（实际计佣金额*佣金比例*最终比例）。如订单完成后发生退款，此金额会更新。
	ActualFee float32 `json:"actualFee"`
	//佣金比例
	CommissionRate float32 `json:"commissionRate"`
	//预估计佣金额，即用户下单的金额(已扣除优惠券、白条、支付优惠、进口税，未扣除红包和京豆)，有时会误扣除运费券金额，完成结算时会在实际计佣金额中更正。如订单完成前发生退款，此金额也会更新。
	EstimateCosPrice float32 `json:"estimateCosPrice"`
	//推客的预估佣金（预估计佣金额*佣金比例*最终比例），如订单完成前发生退款，此金额也会更新。
	EstimateFee float32 `json:"estimateFee"`
	//最终比例（分成比例+补贴比例）
	FinalRate float32 `json:"finalRate"`
	FrozenSkuNum uint64 `json:"frozenSkuNum"`
	Pid string `json:"pid"`
	PositionId uint64 `json:"positionId"`
	//商品单价
	Price float32 `json:"price"`
	SkuId uint64 `json:"skuId"`
	SkuName string `json:"skuName"`
	SkuNum uint64 `json:"skuNum"`
	//商品已退货数量
	SkuReturnNum uint64 `json:"skuReturnNum"`
	//分成比例
	subSideRate float32 `json:"subSideRate"`
	//补贴比例
	SubsidyRate float32 `json:"subsidyRate"`
	//联盟标签数据（整型的二进制字符串(32位)，目前只返回8位：00000001。数据从右向左进行，每一位为1表示符合联盟的标签特征.
	// 第1位：优惠券，第2位：组合推广订单，第3位：拼购订单，第5位：有效首次购订单（00011XXX表示有效首购，最终奖励活动结算金额会结合订单状态判断
	// 以联盟后台对应活动效果数据报表https://union.jd.com/active为准）。
	// 例如：00000001:优惠券，00000010:组合推广订单，00000100:拼购订单，00011000:有效首购，00000111：优惠券+组合推广+拼购等）
	UnionTag string `json:"unionTag"`
	//自定义参数 子联盟ID(需要联系运营开放白名单才能拿到数据)
	SubUnionId string `json:"subUnionId"`
	//订单行维度预估结算时间（格式：yyyyMMdd） ，0：未结算。订单&quot;预估结算时间&quot;仅供参考。账号未通过资质审核或订单发生售后，会影响订单实际结算时间。
	PayMonth string `json:"payMonth"`
	//推客生成推广链接时传入的扩展字段（需要联系运营开放白名单才能拿到数据）。&lt;订单行维度&gt;
	Ext1 string `json:"ext1"`
	//推客的联盟ID
	UnionId string `json:"unionId"`
	//sku维度的有效码
	// （-1：未知,2.无效-拆单,3.无效-取消,4.无效-京东帮帮主订单,5.无效-账号异常,6.无效-赠品类目不返佣
	// 7.无效-校园订单,8.无效-企业订单,9.无效-团购订单,10.无效-开增值税专用发票订单,11.无效-乡村推广员下单
	// 12.无效-自己推广自己下单,13.无效-违规订单,14.无效-来源与备案网址不符,15.待付款,16.已付款,17.已完成,18.已结算（5.9号不再支持结算状态回写展示）
	// ）
	// 注：自2018/7/13起，自己推广自己下单已经允许返佣，故12无效码仅针对历史数据有效
	ValidCode int `json:"validCode"`

}
type OrderResponseResult struct {
	Resp OrderResponse `json:"data"`
}
type UnionOpenOrderQueryResponse struct {
	Result string `json:"result"`
	Code int `json:"code"`
}
type UnionOpenOrderQueryContent struct {
	QueryResp UnionOpenOrderQueryResponse `json:"jd_union_open_order_query_response"`
}

func NewUnionOpenOrderQueryResponse(content []byte) *OrderResponse {
	response := &UnionOpenOrderQueryContent{}
	json.Unmarshal(content, response)
	println(response.QueryResp.Result)
	respResult := &OrderResponse{}
	if response.QueryResp.Code == 0 {
		json.Unmarshal([]byte(response.QueryResp.Result),respResult)
	}

	return respResult
}