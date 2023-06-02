package xpaddle

import "fmt"

type ServerResp struct {
	Msg    string `json:"msg"`
	Status string `json:"status"`
}

func (s *ServerResp) IsSuccess() bool {
	return s.Status == "000"
}
func (s *ServerResp) Error() string {
	return fmt.Sprintf("Status:%s,Msg:%s", s.Status, s.Msg)
}
