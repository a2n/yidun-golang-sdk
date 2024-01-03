package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/pretreatment"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/pretreatment/v1/update"
)

// 提交名单
func main() {

	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := pretreatment.NewPretreatmentClientWithAccessKey("YOUR_SECRET_ID", "YOUR_SECRET_KEY")

	// 设置易盾内容安全分配的businessId
	req := update.NewPretreatmentUpdateRequest("YOUR_BUSSINESS_ID")

	// 设置请求参数

	req.SetId(3770013)
	req.SetEntity("123")
	req.SetBusinessId("YOUR_BUSSINESS_ID")

	response, err := client.Update(req)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		result := response.Result
		resultJson, _ := json.Marshal(result)
		fmt.Println("result: ", string(resultJson))
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
