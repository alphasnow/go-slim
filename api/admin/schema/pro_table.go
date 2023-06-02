package schema

type ProTableReq struct {
	Page int `form:"current"`
	Size int `form:"pageSize"`
}

func (s *ProTableReq) GetPage() int {
	if s.Page <= 0 {
		return 1
	}
	return s.Page
}
func (s *ProTableReq) GetSize() int {
	if s.Size <= 0 {
		return 10
	}
	return s.Size
}

type ProTableRes struct {
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}

// ProTableData 废弃
type ProTableData struct {
	Data  []interface{} `json:"data"`
	Page  int           `json:"page"`
	Total int           `json:"total"`
}

// ProFormOption 废弃
type ProFormOption[V int | string] struct {
	Label string `json:"label"`
	Value V      `json:"value"`
}

// ProTableList 废弃
type ProTableList[M any] struct {
	Data  []M `json:"data"`
	Page  int `json:"page"`
	Total int `json:"total"`
}
