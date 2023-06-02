package schema

import (
	"encoding/json"
)

type Res struct {
	ApiRes
	Errors []string `json:"errors"`
}

func NewRes(code Code, errs ...error) *Res {
	res := &Res{ApiRes: ApiRes{Code: int(code), Message: code.String()}, Errors: []string{}}
	// 解决 errors:null
	if len(errs) > 0 {
		for _, err := range errs {
			if err == nil {
				continue
			}
			res.Errors = append(res.Errors, err.Error())
		}
	}
	return res
}

func (r *Res) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}
