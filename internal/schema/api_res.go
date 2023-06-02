package schema

// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/red-packet-cover/getRedPacketCoverUrl.html
type ApiRes struct {
	Code      int    `json:"err_code"`
	Message   string `json:"err_msg"`
	RequestID string `json:"req_id,omitempty"`
}

type DataRes struct {
	ApiRes
	Data interface{} `json:"data,omitempty"`
}

func NewDataRes(data interface{}) *DataRes {
	resp := &DataRes{Data: map[string]interface{}{}}
	if data != nil {
		resp.Data = data
	}
	return resp
}

//{
//    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ",
//    "token_type": "bearer",
//    "expires_in": 3600
//}
