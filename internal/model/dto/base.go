package dto

type Pager struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

func (item Pager) GetLimit() int {
	if item.PageSize < 1 {
		return 1
	}
	return item.PageSize
}

func (item Pager) GetOffset() int {
	offset := item.PageSize * (item.Page - 1)
	if offset < 0 {
		return 0
	}
	return offset
}
