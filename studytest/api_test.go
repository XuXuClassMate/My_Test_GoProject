package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://8.130.94.143:12345/dolphinscheduler/login?userName=admin&userPassword=dolphinscheduler123"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Host", "8.130.94.143:12345")
	req.Header.Add("language", "zh_CN")
	req.Header.Add("Origin", "http://8.130.94.143:12345")
	req.Header.Add("Referer", "http://8.130.94.143:12345/dolphinscheduler/ui/login")
	//req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36")
	req.Header.Add("sessionId", "")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
