# 滴滴联盟 openAPI go-sdk

引入mod  
```
go get github.com/dunion-openapi-sdk/dunion-go-sdk@master
```
使用方法  
```
c := client.NewUnionClient("appkey", "accesskey")
//日志可选，将在指定目录生成日志
util.InitLogger("./log/union.log")

//或者使用日志注入的方式，需实现两个接口函数：
//Infof(template string, args ...interface{})
//Errorf(template string, args ...interface{})
//然后调用
//util.SetLogger(yourLogger)

//设置全局超时时间
//util.SetTimeoutDuration(2*time.Second)

//或者设置单个接口的超时时间
//link, err := c.GenerateH5Link(context.Background(), 6133, 6834408369283047676, "d", model.Option{Timeout: 2*time.Second})

link, err := c.GenerateH5Link(context.Background(), 6133, 6834408369283047676, "d")
if err != nil {
    fmt.Println(err)
    return
}
```

函数一览  

|  函数原型   | 用途  |
|  ----  | ----  |
| GenerateH5Link(activityID, promotionID int64, sourceID string) (*model.LinkResponse, error) | 生成h5推广链接 |
| GenerateMiniLink(activityID, promotionID int64, sourceID string) (*model.LinkResponse, error) | 生成小程序页面推广路径|  
| GenerateH5Code(dsi, sourceID string) (*model.QrcodeResponse, error)|生成h5二维码，需先取链得到dsi|
| GenerateMiniCode(dsi, sourceID string) (*model.QrcodeResponse, error)|生成小程序太阳码，需先取链得到dsi|
| GeneratePoster(dsi, sourceID string) (*model.PosterResponse, error)|生成推广海报，需先取链得到dsi|
| QueryOrderList(startTime, endTime time.Time, type_ string, page, size int) (*model.OrderResponse, error)|查询订单列表，type_可用枚举见 const.OrderTypeEnergy等|
| MockOrderCallback(dsi string, sourceID string, type_ int) (*model.OrderCallbackResponse, error)|模拟订单回调，需先取链得到 dsi，type_ 可取 consts.MockPay 或 consts.MockRefund; 需在后台配置回调地址｜
| GenerateH5CodeDirectly(activityID, promotionID int64, sourceID string) (*model.QrcodeResponse, error)|直接生成h5推广二维码，会内置请求一次取链接口|
| GenerateMiniCodeDirectly(activityID, promotionID int64, sourceID string) (*model.QrcodeResponse, error)|直接生成小程序推广太阳码，会内置请求一次取链接口|
| GeneratePosterDirectly(activityID, promotionID int64, sourceID string) (*model.PosterResponse, error)|直接生成推广海报，会内置请求一次取链接口|
| SelfQueryOrder(orderID string)(*model.SelfQueryResponse, error)| 订单归因问题自查询|

