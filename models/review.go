package models

type Review struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	BookingID int    `gorm:"not null;index" json:"booking_id"`
	UserID    int    `gorm:"not null;index" json:"user_id"`
	Rating    int    `gorm:"not null" json:"rating"`
	ReviewText string `gorm:"type:text" json:"review_text"`
}

type CreateReviewRequest struct {
	BookingID int    `json:"booking_id" binding:"required"`
	Rating    int    `json:"rating" binding:"required,min=1,max=5"`
	ReviewText string `json:"review_text"`
}

type ReviewResponse struct {
	ID         int    `json:"id"`
	BookingID  int    `json:"booking_id"`
	UserID     int    `json:"user_id"`
	UserName   string `json:"user_name"`
	Rating     int    `json:"rating"`
	ReviewText string `json:"review_text"`
}

