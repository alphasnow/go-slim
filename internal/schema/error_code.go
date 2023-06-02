//go:generate stringer -type=Code -linecomment -output code_string.go

package schema

type Code int

const (
	BadRequest          Code = 10400 // 请求错误
	InternalServerError Code = 10500 // 内部服务错误
	Unauthorized        Code = 10401 // 未授权访问
	NotFound            Code = 10404 // 请求地址错误
	MethodNotAllowed    Code = 10405 // 请求方式错误
	TooManyRequests     Code = 10429 // 请求过快
)

// ServerError
const (
	ServerError         Code = iota + 11000 // 综合服务错误
	ExternalServerError                     // 综合外部服务错误
)

const (
	ReqError Code = iota + 11100 // 请求参数错误
)

const (
	SignJsonError    Code = iota + 11200 // 请求签名错误
	SignJsonParam                        // 请求参数错误
	SignJsonAppid                        // Appid错误
	SignJsonSign                         // 签名验证错误
	SignJsonBody                         // 请求主体错误
	SignJsonNonceStr                     // 随机字符串重复错误
)

const (
	SignUrlError  Code = iota + 11300 // 请求签名错误
	SignUrlParam                      // 请求参数错误
	SignUrlAppid                      // Appid错误
	SignUrlSign                       // 签名验证错误
	SignUrlExpire                     // 请求签名过期
)
