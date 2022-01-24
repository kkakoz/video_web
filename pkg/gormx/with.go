package gormx

import (
	"fmt"
	"gorm.io/gorm"
)

type DBOption func(db *gorm.DB) *gorm.DB

func WithEq(k string, v interface{}) DBOption {
	return func(db *gorm.DB) *gorm.DB {
		switch value := v.(type) {
		case uint8, uint16, uint32, uint64, uint, int8, int16, int32, int64, int:
			if value == 0 {
				return db
			}
		case string:
			if value == "" {
				return db
			}
		}
		return db.Where(fmt.Sprintf("%s = ?", k), v)
	}
}

func WithLike(k, v string) DBOption {
	return func(db *gorm.DB) *gorm.DB {
		if v == "" {
			return db
		}
		return db.Where(fmt.Sprintf("%s like ?", k), "%"+v+"%")
	}
}

func WithWhere(query interface{}, args ...interface{}) DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	}
}

func WithIn(k string, v interface{}) DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s in ?", k), v)
	}
}

// WithPreload preload的model的字段不可以作为分页条件
func WithPreload(name string, args ...interface{}) DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(name, args...)
	}
}

// WithJoins 分页条件最好加在where里面,如果是多表关联的接口最好自己单独写
func WithJoins(name string, args ...interface{}) DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins(name, args...)
	}
}
