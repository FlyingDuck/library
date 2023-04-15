package model

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"
)

type Book struct {
	ID        int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	BookName  string `gorm:"column:book_name" json:"book_name"`
	Author    string `gorm:"column:author" json:"author"`
	Category  string `gorm:"column:category" json:"category"`
	Location  string `gorm:"column:location" json:"location"`
	Notice    string `gorm:"column:notice" json:"notice"`
	Source    string `gorm:"column:source" json:"source"`
	State     string `gorm:"column:state" json:"state"`
	Keywords  CSSA   `gorm:"column:keywords" json:"keywords"`
	Images    CSSA   `gorm:"column:images" json:"images"`
	Abstract  string `gorm:"column:abstract" json:"abstract"`
	Theme     string `gorm:"column:theme" json:"theme"`
	Publisher string `gorm:"column:publisher" json:"publisher"`
	ISBN      string `gorm:"column:isbn" json:"isbn"`

	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"-"` // 创建时间
	ModifyTime time.Time `gorm:"column:modify_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"-"` // 更新时间
}

func (b *Book) TableName() string {
	return "book"
}

type CSSA []string // Comma Separated String Array

func (m CSSA) Value() (driver.Value, error) {
	if len(m) == 0 {
		return "", nil
	}
	return strings.Join(m, ","), nil
}

func (m *CSSA) Scan(src interface{}) error {
	if src == nil {
		*m = CSSA{}
		return nil
	}
	if sv, err := driver.String.ConvertValue(src); err == nil {
		if v, ok := sv.(string); ok {
			*m = strings.Split(v, ",")
			return nil
		}
		if v, ok := sv.([]byte); ok {
			*m = strings.Split(string(v), ",")
			return nil
		}
	}
	return errors.New("failed to scan CSSA(Comma Separated String Array)")
}
