package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	http "gopkg.in/gavv/httpexpect.v2"
	"testing"
)

func TestQueryBuildVersion(t *testing.T) {
	baseurl := http.New(t, "http://ws4:12345/dolphinscheduler")

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		//DisableColors: true,
		FullTimestamp: true,
	})

	// 发送查询版本请求
	resp := baseurl.GET("/version/queryBuildVersion").
		WithHeader("Accept", "application/json, text/plain, */*").
		WithHeader("sessionId", "da9e7f28-3647-4c55-a0c0-8f9b657a0029").
		WithHeader("Referer", "http://ws4:12345/dolphinscheduler/ui/home").
		//WithHeader("language", "zh_CN").
		//WithHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36").
		Expect().
		Status(200). // 请根据实际情况调整期望的状态码
		JSON().
		Object()

	// 在这里添加验证接口返回中data等于文本类型的"2.4"的断言
	data := resp.Value("data").String().Raw()
	logger.Info("response：", resp)
	if data != "2.4" {
		t.Errorf("t.Errorf Expected data to be '2.4', but got '%s'", data)
		logger.Errorf("logger.Errorf Expected data to be '2.4', but got '%s'", data)

	}
	data1 := resp.Raw()

	// 打印响应原始数据
	fmt.Println("response：", data1)

}

func main() {
	t := testing.T{}
	TestQueryBuildVersion(&t) // Pass the address of the variable
}
