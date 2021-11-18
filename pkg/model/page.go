package model

import "gorm.io/gorm"

type Pager struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func (p *Pager) Paging(db *gorm.DB) *gorm.DB{
	if p.Page >= 0 {
		db = db.Offset(p.Page * p.GetSize())
	}
	if p.PageSize < 1 {
		db = db.Limit(10)
	}
	return db
}

func (p *Pager) GetSize() int {
	if p.PageSize < 1 {
		return 10
	}
	return p.PageSize
}
