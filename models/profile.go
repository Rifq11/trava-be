package models

type ProfileResponse struct {
	User    User        `json:"user"`
	Profile interface{} `json:"profile"` // Can be UserProfile or AdminProfile
}

type CompleteProfileRequest struct {
	Phone     *string `json:"phone"`
	Address   *string `json:"address"`
	BirthDate *string `json:"birth_date"`
	UserPhoto *string `json:"user_photo"`
	IsAdmin   *bool   `json:"is_admin"`
}

type AdminProfile struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      int    `gorm:"not null;index" json:"user_id"`
	Phone       string `gorm:"type:varchar(50)" json:"phone"`
	Address     string `gorm:"type:text" json:"address"`
	BirthDate   string `gorm:"type:date" json:"birth_date"`
	UserPhoto   string `gorm:"type:varchar(500)" json:"user_photo"`
	IsCompleted bool   `gorm:"default:false" json:"is_completed"`
}

