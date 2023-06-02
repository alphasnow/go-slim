package xsignjson

import "strings"

// https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=4_3
// X-Appid
// X-Nonce-Str
// X-Sign

type ISign interface {
	Signature(interface{}) (string, error)
	Verify(data map[string]interface{}) bool
}

var _ ISign = (*Sign)(nil)

type Sign struct {
	Appid    string `form:"appid" json:"appid" header:"X-Appid" binding:"required,max=32" validate:"required,min=8,max=32"`
	NonceStr string `form:"nonce_str" json:"nonce_str" header:"X-Nonce-Str" binding:"required,max=32" validate:"required,min=8,max=32"`
	Sign     string `form:"sign" json:"sign" header:"X-Sign" binding:"required,max=64" validate:"required,min=32,max=64"`
	// SignType string `form:"sign_type" json:"sign_type,omitempty" header:"X-Sign-Type" binding:"max=12" validate:"max=12"`
	SignType string `json:"-"`
	Appkey   string `json:"-"`
}

func (s *Sign) getCommonMap() map[string]string {
	res := make(map[string]string)
	res["appid"] = s.Appid
	res["nonce_str"] = s.NonceStr
	if s.SignType != "" {
		res["sign_type"] = s.SignType
	}
	if s.Sign != "" {
		res["sign"] = s.Sign
	}
	return res
}

func (s *Sign) encrypt(data string) string {
	var reqSign string
	switch s.SignType {
	case "hash_hmac":
		reqSign = xHmacSHA256(data, s.Appkey)
	default:
		reqSign = xMD5(data)
	}
	return reqSign
}

// Signature 生成签名
// 上传文件使用Base64
// 传复杂数据使用JSON序列化的字符串
func (s *Sign) Signature(data interface{}) (string, error) {
	// 数据转字符串
	resMapStr, err := jsonStructToStringMap(data)
	if err != nil {
		return "", err
	}

	s.Sign = s.generate(resMapStr)
	return s.Sign, nil
}

func (s *Sign) generate(data map[string]string) string {
	// 附加公共字段
	for k, v := range s.getCommonMap() {
		data[k] = v
	}

	// 排除sign
	if _, ok := data["sign"]; ok {
		delete(data, "sign")
	}

	// 字符串排序
	resSortStr := stringMapToSortMap(data)
	// 附加秘钥
	resSortStr += "&key=" + s.Appkey
	// 字符串加密
	reqSign := s.encrypt(resSortStr)
	// 字符串大写
	reqSign = strings.ToUpper(reqSign)
	return reqSign
}

func (s *Sign) Verify(data map[string]interface{}) bool {
	// 检测
	if s.Appkey == "" {
		return false
	}
	reqSign, ok := data["sign"]
	if ok == false || reqSign == "" {
		return false
	}

	// 数据转字符串
	reqMapStr, err := jsonMapToStringMap(data)
	if err != nil {
		return false
	}

	reqSign = s.generate(reqMapStr)
	return reqSign == s.Sign
}

func (s *Sign) Request(data map[string]interface{}) (map[string]interface{}, error) {
	// 数据转字符串
	resMapStr, err := jsonMapToStringMap(data)
	if err != nil {
		return nil, err
	}

	s.Sign = s.generate(resMapStr)

	// 附加公共字段
	for k, v := range s.getCommonMap() {
		data[k] = v
	}
	return data, nil
}
