package types

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

type PostJsonRequest struct {
	*BaseRequestConstruct
	Version         string
	Timestamp       int64
	Nonce           string
	SignatureMethod auth.SignatureMethod
	GzipCompress    bool
}

func NewPostJsonRequest() *PostJsonRequest {
	result := &PostJsonRequest{
		BaseRequestConstruct: NewBaseRequestConstruct(),
		Timestamp:            time.Now().UnixNano() / int64(time.Millisecond),
		Nonce:                RandUUID(),
		GzipCompress:         false,
	}
	result.Method = http.HttpMethodPost
	return result
}

// GetSignParams 获取待签名参数列表
func (r *PostJsonRequest) GetSignParams() map[string]string {
	// 创建一个新map，将原始map的所有键值对复制到新map中，避免修改原始map
	params := make(map[string]string)
	for k, v := range r.CustomParams {
		params[k] = v
	}
	// 修改新map的值
	params["version"] = r.Version
	params["timestamp"] = strconv.FormatInt(r.Timestamp, 10)
	params["nonce"] = r.Nonce
	if r.SignatureMethod != "" {
		params["signatureMethod"] = string(r.SignatureMethod)
	}
	return params
}

func (r *PostJsonRequest) GetHeaders() map[string]string {
	headers := make(map[string]string)
	for k, v := range r.BaseRequestConstruct.GetHeaders() {
		headers[k] = v
	}
	headers[http.ContentType] = "application/json;charset=utf-8"
	if r.GzipCompress {
		headers[http.ContentEncoding] = "gzip"
	}
	return headers
}

// GetBodyWithSign 构建 body，将参数以json格式组装放入 body 中
func (r *PostJsonRequest) GetBodyWithSign(signer auth.Signer, credentials auth.Credentials) ([]byte, error) {
	params := r.GetSignParams()
	signResult := signer.GenSignature(credentials, params)
	params["secretId"] = signResult.SecretId
	params["signature"] = signResult.Signature
	for k, v := range r.NonSignParams {
		params[k] = v
	}

	// 将map[string]interface{}类型的数据转换为JSON格式的字节数组
	jsonBytes, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	if r.GzipCompress {
		body, err := ToGzipBytes(jsonBytes)
		if err == nil {
			return body, nil
		} else {
			return jsonBytes, nil
		}
	}
	return jsonBytes, nil
}

// ToHttpRequest 构建 http 请求
func (r *PostJsonRequest) ToHttpRequest(signer auth.Signer, credentials auth.Credentials) (http.Request, error) {
	body, err := r.GetBodyWithSign(signer, credentials)
	if err != nil {
		return nil, err
	}
	req := http.HttpRequest{
		MethodValue:  string(r.Method),
		UrlValue:     r.AssembleUrl(),
		HeadersValue: r.GetHeaders(),
		BodyValue:    body,
	}
	return &req, nil
}

func (p *PostJsonRequest) SetVersion(version string) {
	p.Version = version
}

func (p *PostJsonRequest) SetTimestamp(timestamp int64) {
	p.Timestamp = timestamp
}

func (p *PostJsonRequest) SetNonce(nonce string) {
	p.Nonce = nonce
}

func (p *PostJsonRequest) SetSignatureMethod(signatureMethod auth.SignatureMethod) {
	p.SignatureMethod = signatureMethod
}

func (p *PostJsonRequest) SetGzipCompress(gzipCompress bool) {
	p.GzipCompress = gzipCompress
}
