package request

type Pager struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

func (item Pager) GetLimit() int {
	return item.PageSize
}

func (item Pager) GetOffset() int {
	return item.PageSize * item.Page
}
