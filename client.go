package client

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dunion-openapi-sdk/dunion-go-sdk/const"
	"github.com/dunion-openapi-sdk/dunion-go-sdk/model"
	"github.com/dunion-openapi-sdk/dunion-go-sdk/util"
	"time"
)

type client struct {
	AppKey    string
	AccessKey string
}

type UnionClient interface {
	GenerateH5Link(ctx context.Context, activityID, promotionID int64, sourceID string, opt ...model.Option) (*model.LinkResponse, error)
	GenerateMiniLink(ctx context.Context, activityID, promotionID int64, sourceID string, opt ...model.Option) (*model.LinkResponse, error)
	GenerateH5Code(ctx context.Context, dsi, sourceID string, opt ...model.Option) (*model.QrcodeResponse, error)
	GenerateMiniCode(ctx context.Context, dsi, sourceID string, opt ...model.Option) (*model.QrcodeResponse, error)
	GeneratePoster(ctx context.Context, dsi, sourceID string, opt ...model.Option) (*model.PosterResponse, error)
	QueryOrderList(ctx context.Context, startTime, endTime time.Time, type_ string, page, size int, opt ...model.Option) (*model.OrderResponse, error)
	GenerateH5CodeDirectly(ctx context.Context, activityID, promotionID int64, sourceID string, opt ...model.Option) (*model.QrcodeResponse, error)
	MockOrderCallback(ctx context.Context, dsi string, sourceID string, type_ int, opt ...model.Option) (*model.OrderCallbackResponse, error)
	GenerateMiniCodeDirectly(ctx context.Context, activityID, promotionID int64, sourceID string, opt ...model.Option) (*model.QrcodeResponse, error)
	GeneratePosterDirectly(ctx context.Context, activityID, promotionID int64, sourceID string, opt ...model.Option) (*model.PosterResponse, error)
	SelfQueryOrder(ctx context.Context, orderID string, opt ...model.Option) (*model.SelfQueryResponse, error)
}

func NewUnionClient(appKey string, accessKey string) UnionClient {
	return &client{
		AppKey:    appKey,
		AccessKey: accessKey,
	}
}

//GenerateH5Link 生成h5推广链接
func (s client) GenerateH5Link(ctx context.Context, activityID, promotionID int64, sourceID string, opt ...model.Option) (*model.LinkResponse, error) {
	body := map[string]interface{}{
		"activity_id":  activityID,
		"link_type":    "h5",
		"promotion_id": promotionID,
		"source_id":    sourceID,
	}
	response, err := util.Post(ctx, s.AppKey, s.AccessKey, consts.GenerateLinkUrl, body, opt...)
	if err != nil {
		return nil, err
	}
	result := &model.LinkResponse{}
	err = json.Unmarshal(response, result)
	return result, err
}

//GenerateMiniLink 生成小程序页面推广路径
func (s client) GenerateMiniLink(ctx context.Context, activityID, promotionID int64, sourceID string, opt ...model.Option) (*model.LinkResponse, error) {
	body := map[string]interface{}{
		"activity_id":  activityID,
		"link_type":    "mini",
		"promotion_id": promotionID,
		"source_id":    sourceID,
	}
	response, err := util.Post(ctx, s.AppKey, s.AccessKey, consts.GenerateLinkUrl, body, opt...)
	if err != nil {
		return nil, err
	}
	result := &model.LinkResponse{}
	err = json.Unmarshal(response, result)
	return result, err
}

//GenerateH5Code 生成h5二维码，需先取链得到dsi
func (s client) GenerateH5Code(ctx context.Context, dsi, sourceID string, opt ...model.Option) (*model.QrcodeResponse, error) {
	param := map[string]interface{}{
		"dsi":       dsi,
		"source_id": sourceID,
		"type":      "h5",
	}
	response, err := util.Get(ctx, s.AppKey, s.AccessKey, consts.GenerateQrCodeUrl, param, opt...)
	if err != nil {
		return nil, err
	}
	result := &model.QrcodeResponse{}
	err = json.Unmarshal(response, result)
	return result, err
}

//GenerateMiniCode 生成小程序太阳码，需先取链得到dsi
func (s client) GenerateMiniCode(ctx context.Context, dsi, sourceID string, opt ...model.Option) (*model.QrcodeResponse, error) {
	param := map[string]interface{}{
		"dsi":       dsi,
		"source_id": sourceID,
		"type":      "mini",
	}
	response, err := util.Get(ctx, s.AppKey, s.AccessKey, consts.GenerateQrCodeUrl, param, opt...)
	if err != nil {
		return nil, err
	}
	result := &model.QrcodeResponse{}
	err = json.Unmarshal(response, result)
	return result, err
}

//GeneratePoster 生成推广海报，需先取链得到dsi
func (s client) GeneratePoster(ctx context.Context, dsi, sourceID string, opt ...model.Option) (*model.PosterResponse, error) {
	param := map[string]interface{}{
		"dsi":       dsi,
		"source_id": sourceID,
	}
	response, err := util.Get(ctx, s.AppKey, s.AccessKey, consts.GeneratePosterUrl, param, opt...)
	if err != nil {
		return nil, err
	}
	result := &model.PosterResponse{}
	err = json.Unmarshal(response, result)
	return result, err
}

//QueryOrderList 查询订单列表，type_可用枚举见 const.OrderTypeEnergy等
func (s client) QueryOrderList(ctx context.Context, startTime, endTime time.Time, type_ string, page, size int, opt ...model.Option) (*model.OrderResponse, error) {
	if page < 0 || page > 100 || size < 0 || size > 100 {
		return nil, errors.New("分页参数不合法")
	}
	param := map[string]interface{}{
		"pay_start_time": startTime.Unix(),
		"pay_end_time":   endTime.Unix(),
		"page":           page,
		"size":           size,
	}
	if len(type_) > 0 {
		param["type"] = type_
	}
	response, err := util.Get(ctx, s.AppKey, s.AccessKey, consts.QueryOrderUrl, param, opt...)
	if err != nil {
		return nil, err
	}
	result := &model.OrderResponse{}
	err = json.Unmarshal(response, result)
	return result, err
}

//MockOrderCallback 模拟订单回调，需先取链得到 dsi，type_ 可取 consts.MockPay 或 consts.MockRefund; 需在后台配置回调地址
func (s client) MockOrderCallback(ctx context.Context, dsi string, sourceID string, type_ int, opt ...model.Option) (*model.OrderCallbackResponse, error) {
	param := map[string]interface{}{
		"dsi":       dsi,
		"source_id": sourceID,
		"type":      type_,
	}
	response, err := util.Get(ctx, s.AppKey, s.AccessKey, consts.MockOrderUrl, param, opt...)
	if err != nil {
		return nil, err
	}
	result := &model.OrderCallbackResponse{}
	err = json.Unmarshal(response, result)
	return result, err
}

//GenerateH5CodeDirectly 直接生成h5推广二维码，会内置请求一次取链接口
func (s client) GenerateH5CodeDirectly(ctx context.Context, activityID, promotionID int64, sourceID string, opt ...model.Option) (*model.QrcodeResponse, error) {
	link, err := s.GenerateMiniLink(ctx, activityID, promotionID, sourceID, opt...)
	if err != nil {
		return nil, err
	}
	dsi := link.Data.DSI
	return s.GenerateH5Code(ctx, dsi, sourceID, opt...)
}

//GenerateMiniCodeDirectly 直接生成小程序推广太阳码，会内置请求一次取链接口
func (s client) GenerateMiniCodeDirectly(ctx context.Context, activityID, promotionID int64, sourceID string, opt ...model.Option) (*model.QrcodeResponse, error) {
	link, err := s.GenerateMiniLink(ctx, activityID, promotionID, sourceID, opt...)
	if err != nil {
		return nil, err
	}
	dsi := link.Data.DSI
	return s.GenerateMiniCode(ctx, dsi, sourceID, opt...)
}

//GeneratePosterDirectly 直接生成推广海报，会内置请求一次取链接口
func (s client) GeneratePosterDirectly(ctx context.Context, activityID, promotionID int64, sourceID string, opt ...model.Option) (*model.PosterResponse, error) {
	link, err := s.GenerateMiniLink(ctx, activityID, promotionID, sourceID, opt...)
	if err != nil {
		return nil, err
	}
	dsi := link.Data.DSI
	return s.GeneratePoster(ctx, dsi, sourceID, opt...)
}

//SelfQueryOrder 订单归因问题自查询
func (s client) SelfQueryOrder(ctx context.Context, orderID string, opt ...model.Option) (*model.SelfQueryResponse, error) {
	param := map[string]interface{}{
		"order_id": orderID,
	}
	response, err := util.Get(ctx, s.AppKey, s.AccessKey, consts.SelfQueryUrl, param, opt...)
	if err != nil {
		return nil, err
	}
	result := &model.SelfQueryResponse{}
	err = json.Unmarshal(response, result)
	return result, err
}
