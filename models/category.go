package models

// 文章分类

type Category struct {
	ID       uint8     `json:"id" gorm:"primary_key"`
	Name     string    `json:"name" gorm:"not null;type:varchar(50);unique"`
	CreatedAt Time   `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt Time   `json:"updated_at" gorm:"type:timestamp"`

}
