package dto

type Pager struct {
	Page     int `json:"page" query:"page"`
	PageSize int `json:"page_size" query:"page_size"`
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
		return 15
	}
	if offset > 200 {
		return 200
	}
	return offset
}

type ID struct {
	ID int64 `json:"id"`
}

type LastId struct {
	LastId int64 `json:"last_id" query:"last_id"`
}
