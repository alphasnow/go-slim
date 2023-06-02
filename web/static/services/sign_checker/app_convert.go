package sign_checker

import (
	"fmt"
	"strconv"
)

type AppCache struct {
	ID     uint
	Appkey string
	Appid  string
}

type AppConvert struct {
}

func (s *AppConvert) ToMapString(app AppCache) map[string]string {
	res := map[string]string{}
	res["id"] = fmt.Sprintf("%d", app.ID)
	res["appid"] = app.Appid
	res["appkey"] = app.Appkey
	return res
}

func (s *AppConvert) ToAppModel(data map[string]string) AppCache {
	res := AppCache{}
	id, _ := strconv.Atoi(data["id"])
	res.ID = uint(id)
	res.Appid = data["appid"]
	res.Appkey = data["appkey"]
	return res
}
