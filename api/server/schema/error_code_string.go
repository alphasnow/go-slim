// Code generated by "stringer -type=Code -linecomment -output code_string.go"; DO NOT EDIT.

package schema

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BadRequest-10400]
	_ = x[InternalServerError-10500]
	_ = x[Unauthorized-10401]
	_ = x[NotFound-10404]
	_ = x[MethodNotAllowed-10405]
	_ = x[TooManyRequests-10429]
	_ = x[ServerError-11000]
	_ = x[ExternalServerError-11001]
	_ = x[ReqError-11100]
	_ = x[SignJsonError-11200]
	_ = x[SignJsonParam-11201]
	_ = x[SignJsonAppid-11202]
	_ = x[SignJsonSign-11203]
	_ = x[SignJsonBody-11204]
	_ = x[SignJsonNonceStr-11205]
	_ = x[SignUrlError-11300]
	_ = x[SignUrlParam-11301]
	_ = x[SignUrlAppid-11302]
	_ = x[SignUrlSign-11303]
	_ = x[SignUrlExpire-11304]
}

const (
	_Code_name_0 = "请求错误未授权访问"
	_Code_name_1 = "请求地址错误请求方式错误"
	_Code_name_2 = "请求过快"
	_Code_name_3 = "内部服务错误"
	_Code_name_4 = "综合服务错误综合外部服务错误"
	_Code_name_5 = "请求参数错误"
	_Code_name_6 = "请求签名错误请求参数错误Appid错误签名验证错误请求主体错误随机字符串重复错误"
	_Code_name_7 = "请求签名错误请求参数错误Appid错误签名验证错误请求签名过期"
)

var (
	_Code_index_0 = [...]uint8{0, 12, 27}
	_Code_index_1 = [...]uint8{0, 18, 36}
	_Code_index_4 = [...]uint8{0, 18, 42}
	_Code_index_6 = [...]uint8{0, 18, 36, 47, 65, 83, 110}
	_Code_index_7 = [...]uint8{0, 18, 36, 47, 65, 83}
)

func (i Code) String() string {
	switch {
	case 10400 <= i && i <= 10401:
		i -= 10400
		return _Code_name_0[_Code_index_0[i]:_Code_index_0[i+1]]
	case 10404 <= i && i <= 10405:
		i -= 10404
		return _Code_name_1[_Code_index_1[i]:_Code_index_1[i+1]]
	case i == 10429:
		return _Code_name_2
	case i == 10500:
		return _Code_name_3
	case 11000 <= i && i <= 11001:
		i -= 11000
		return _Code_name_4[_Code_index_4[i]:_Code_index_4[i+1]]
	case i == 11100:
		return _Code_name_5
	case 11200 <= i && i <= 11205:
		i -= 11200
		return _Code_name_6[_Code_index_6[i]:_Code_index_6[i+1]]
	case 11300 <= i && i <= 11304:
		i -= 11300
		return _Code_name_7[_Code_index_7[i]:_Code_index_7[i+1]]
	default:
		return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}