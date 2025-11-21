package models

import "../Models/time"

type ActivityLog struct {
	ID            int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        int       `gorm:"column:user_id;not null;index" json:"user_id"`
	DestinationID *int      `gorm:"column:destination_id;index" json:"destination_id"`
	ActivityType  string    `gorm:"column:activity_type;type:varchar(255);not null" json:"activity_type"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (ActivityLog) TableName() string {
	return "user_activity_log"
}

type LogActivityRequest struct {
	DestinationID *int   `json:"destination_id"`
	ActivityType  string `json:"activity_type" binding:"required"`
}
