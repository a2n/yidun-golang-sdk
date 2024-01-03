package submit

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

type FileSubmitV2Response struct {
	*types.CommonResponse
	Result *FileSubmitResult `json:"result"`
}

type FileSubmitResult struct {
	//本次请求数据标识，可以根据该标识查询数据最新结果
	TaskId string `json:"taskId,omitempty"`
	//数据唯一标识，能够根据该值定位到该条数据，如对数据结果有异议，可以发送该值给客户经理查询
	DataId string `json:"dataId,omitempty"`
	//缓冲池当前缓冲数量
	DealingCount int64 `json:"dealingCount,omitempty"`
}
