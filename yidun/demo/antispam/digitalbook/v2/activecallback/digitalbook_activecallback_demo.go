package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yidun/yidun-golang-sdk/yidun/core/util"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/digitalbook/v2/callback"
)

const (
	SECRET_KEY = "YOUR_SECRET_KEY"
)

/**
 * 接收主动回调请求
 */
func main() {
	http.HandleFunc("/callback/receive/digitalbook", handleCallback) //设置主动回调url

	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 处理回调数据
func handleCallback(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数，默认是不会解析的
	// 获取签名参数和其他参数
	params := make(map[string]string)
	for k, v := range r.Form {
		if len(v) > 0 {
			params[k] = v[0]
		}
	}
	verifyResult := util.VerifySignature(r.Form, SECRET_KEY)
	//签名验证通过
	if verifyResult {
		callbackRequest := callback.NewDigitalBookActiveCallbackV2Request(params)
		// 解析得到数字阅读检测结果
		// 根据需要解析字段，具体返回字段的说明，请参考官方接口数字阅读中字段说明 https://support.dun.163.com/documents/601492044329328640?docId=601495324846338048
		result := &callback.DigitalBookCheckResult{}
		json.Unmarshal([]byte(*callbackRequest.CallbackData), result)

		if result != nil {
			// do something
		}
	}

}
