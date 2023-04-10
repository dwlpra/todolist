package entity

import "time"

type Todo struct {
	ID              uint64    `gorm:"primaryKey" json:"id"`
	ActivityGroupID uint64    `gorm:"index" json:"activity_group_id"`
	Title           string    `gorm:"size:255" json:"title"`
	IsActive        bool      `gorm:"type:TINYINT(1);default:1" json:"is_active"`
	Priority        string    `gorm:"default:'very-high'" json:"priority"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
