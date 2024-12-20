package model

// swagger:model
type ResponseOrderItem struct {
	// 标题
	Title string `json:"title"`
	// 订单id
	OrderId string `json:"order_id"`
	// 业务线 `159`: 滴滴加油<br>
	// `210`: 滴滴网约车<br>
	// `393`: 滴滴货运<br>
	// `500`: 花小猪<br>
	// `120`: 滴滴代驾
	ProductId string `json:"product_id"`
	// 支付时间
	PayTime int64 `json:"pay_time"`
	// 支付金额，单位：分
	PayPrice int64 `json:"pay_price"`
	// CPA类型<br>
	// `cpa_normal`: 普通CPA（新用户奖励）
	RefundPrice int64 `json:"refund_price"`
	// 退款时间，秒级时间戳
	RefundTime int64 `json:"refund_time"`
	// CPS返佣金额，单位：分
	CpsProfit int64 `json:"cps_profit"`
	// CPA返佣金额，单位：分
	CpaProfit int64 `json:"cpa_profit"`
	// CPA类型
	CpaType string `json:"cpa_type"`
	// 推送状态: 1.已预估归因  2.预估订单已推送 3.预估订单推送失败  4.结算已提交  5.结算提交中 6.结算取消  7.结算成功  8.结算失败
	Status int `json:"status"`
	// 推广位ID
	PromotionId int `json:"promotion_id"`
	// 来源ID
	SourceId string `json:"source_id"`
	// 是否被风控
	IsRisk int `json:"is_risk"`
	// 下单用户openUID
	// example:ecca7d66c984706aa94a15b656db2538
	OpenUID string `json:"open_uid" structs:"open_uid"`
	// 订单状态 `2`:已付款
	// `8`:已退款
	// example:2
	OrderStatus int `json:"order_status" structs:"order_status"`
}

//swagger:model
type OrderResponse struct {
	Response
	Data struct {
		// 总数量
		// example: 1
		Total     int                  `json:"total"`
		OrderList []*ResponseOrderItem `json:"order_list"`
	} `json:"data"`
}

//swagger:model
type OrderCallbackResponse struct {
	Response
	Data struct {
		// 标题
		Title string `json:"title"`
		// 订单id
		OrderId string `json:"order_id"`
		// 业务线 `159`: 滴滴加油<br>
		// `210`: 滴滴网约车<br>
		// `393`: 滴滴货运<br>
		// `500`: 花小猪<br>
		// `120`: 滴滴代驾
		ProductId string `json:"product_id"`
		// 支付时间
		PayTime int64 `json:"pay_time"`
		// 支付金额，单位：分
		PayPrice int64 `json:"pay_price"`
		// 退款金额，单位：分
		RefundPrice int64 `json:"refund_price"`
		// 退款时间，秒级时间戳
		RefundTime int64 `json:"refund_time"`
		// CPS返佣金额，单位：分
		CpsProfit int64 `json:"cps_profit"`
		// CPA返佣金额，单位：分
		CpaProfit int64 `json:"cpa_profit"`
		// CPA类型
		CpaType string `json:"cpa_type"`
		// 推送状态: 1.已预估归因  2.预估订单已推送 3.预估订单推送失败  4.结算已提交  5.结算提交中 6.结算取消  7.结算成功  8.结算失败
		Status int `json:"status"`
		// 推广位ID
		PromotionId string `json:"promotion_id"`
		// 来源ID
		SourceId string `json:"source_id"`
		// 是否被风控
		IsRisk int `json:"is_risk"`
	} `json:"data"`
}

type SelfQueryResponse struct {
	Response
	Data ResponseOrderSelfQuery `json:"data"`
}

//swagger:model
type ResponseOrderSelfQuery struct {
	// 推广成功列表
	EstimateSuccessList []*EstimateSuccessData `json:"estimate_success_list,omitempty" structs:"estimate_success_list"`
	// 推广失败列表
	EstimateFailList []*EstimateFailData `json:"estimate_fail_list,omitempty" structs:"estimate_fail_list"`
}

//swagger:model
type EstimateSuccessData struct {
	// 推广成功时间
	// example:2022-01-10 10:30:00
	EstimateTime string `json:"estimate_time" structs:"estimate_time"`
	// 推广渠道
	// example:当前登陆账号注册的公司名
	EstimateChannel string `json:"estimate_channel" structs:"estimate_channel"`
	// 领券状态 `1`:成功
	// `2`:失败
	// example:1
	ReceiveStatus int `json:"receive_status" structs:"receive_status"`
	// 领券时间
	// example:2022-01-10 10:00:00
	ReceiveTime string `json:"receive_time" structs:"receive_time"`
	// 业务线名称
	// example:网约车
	SceneName string `json:"scene_name" structs:"scene_name"`
}

//swagger:model
type EstimateFailData struct {
	// 失败原因
	// example:未查询到有效绑定关系
	FailReason string `json:"fail_reason" structs:"fail_reason"`
	// 业务线名称
	// example:网约车
	SceneName string `json:"scene_name" structs:"scene_name"`
}
