package xsignjson

type SignReq struct {
	Sign    Sign        `json:",inline"`
	Payload interface{} `json:",inline"`
}
