package model

import "time"

type Model struct {
	ID         uint32     `gorm:"primary_key" json:"id"`
	CreatedBy  string     `json:"created_by"`
	ModifiedBy string     `json:"modified_by"`
	CreatedOn  time.Time  `gorm:"type:datetime" json:"created_on"`
	ModifiedOn time.Time  `gorm:"type:datetime" json:"modified_on"`
	DeletedOn  *time.Time `gorm:"type:datetime" json:"deleted_on"`
	IsDel      uint8      `gorm:"index" json:"is_del"`
}

// gorm的Model
// type Model struct {
// 	ID        uint `gorm:"primarykey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt DeletedAt `gorm:"index"`
// }
