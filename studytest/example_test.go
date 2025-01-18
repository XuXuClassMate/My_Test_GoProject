package main

import (
	"testing"

	http "gopkg.in/gavv/httpexpect.v2"
)

type Response http.Response

func TestAPI(t *testing.T) {
	// 创建一个 httpexpect 实例
	login := http.New(t, "http://www.xuxuclassmate.cn:12345/dolphinscheduler")

	// 发送 GET 请求，并断言响应
	login.POST("/login").
		WithQuery("userName", "admin").
		WithQuery("userPassword", "dolphinscheduler123").Expect().Status(200).JSON().String()

}
