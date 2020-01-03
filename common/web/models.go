package web

type Result struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type BsmResult struct {
	Success  bool
	Message  string
	Data     interface{}
	Status   string
	Solution string
}

type PageParam struct {
	Rows  int                    `json:"rows"`
	Page  int                    `json:"page"`
	Param map[string]interface{} `json:"param"`
}

func DefaultPageParam(p *PageParam) {
	if &p == nil {
		p = new(PageParam)
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Rows < 1 {
		p.Rows = 1
	}
	if p.Param == nil {
		p.Param = map[string]interface{}{}
	}
}
