package main

import (
	"context"
	"fmt"
	"github.com/dunion-openapi-sdk/dunion-go-sdk/client"
	"github.com/dunion-openapi-sdk/dunion-go-sdk/util"
)

func main() {
	c := client.NewUnionClient("Fyl1RBOOqAKp1a1s", "CXmIHZd52ks1EjxQOUalYLcbvQEYl3f9")
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
	//link, err := c.GenerateH5Link(context.Background(), 6133, 6834408369283047676, "d", model.Option{Timeout: time.Millisecond})

	link, err := c.GenerateH5Link(context.Background(), 207811824611, 7275712098601201665, "d")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("返回参数：%+v", link)
	fmt.Println(link.Data.Link)
}
