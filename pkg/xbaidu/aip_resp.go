package xbaidu

type bodySegResp struct {
	LogId      int64  `json:"log_id"`
	Labelmap   string `json:"labelmap"`
	Scoremap   string `json:"scoremap"`
	Foreground string `json:"foreground"`
	PersonNum  int    `json:"person_num"`
	PersonInfo []struct {
		Height float64 `json:"height"`
		Width  float64 `json:"width"`
		Top    float64 `json:"top"`
		Score  float64 `json:"score"`
		Left   float64 `json:"left"`
	} `json:"person_info"`
}
