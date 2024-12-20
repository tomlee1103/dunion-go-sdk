package consts

const (
	GenerateLinkUrl   = "https://union.didi.cn/openapi/v1.0/link/generate"
	GenerateQrCodeUrl = "https://union.didi.cn/openapi/v1.0/code/generate"
	GeneratePosterUrl = "https://union.didi.cn/openapi/v1.0/poster/generate"
	QueryOrderUrl     = "https://union.didi.cn/openapi/v1.0/order/list"
	MockOrderUrl      = "https://union.didi.cn/openapi/v1.0/orderMock/callback"
	SelfQueryUrl      = "https://union.didi.cn/openapi/v1.0/order/selfQuery"
)

//订单查询type枚举值
const (
	OrderTypeAll     = ""            //全部
	OrderTypeEnergy  = "energy"      //滴滴加油
	OrderTypeCar     = "online_car"  //网约车
	OrderTypeFreight = "freight"     //货运
	OrderTypeHxz     = "king_flower" //花小猪
	OrderTypeDaijia  = "daijia"      //代驾
)

//mock订单回调类型枚举值
const (
	MockPay    = 0 //支付
	MockRefund = 1 //退款
)

const (
	UserAgent  = "User-Agent"
	SDKVersion = "dunion-go-openapi-sdk-1.0"
	TraceID    = "Didi-Header-Rid"
)
