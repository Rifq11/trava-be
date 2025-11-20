package models

type Payment struct {
	ID           int    `gorm:"primaryKey;autoIncrement" json:"id"`
	BookingID    int    `gorm:"not null;index" json:"booking_id"`
	Amount       int    `gorm:"not null" json:"amount"`
	PaymentStatus string `gorm:"type:varchar(255);not null" json:"payment_status"`
}

type CreatePaymentRequest struct {
	BookingID int `json:"booking_id" binding:"required"`
	Amount    int `json:"amount" binding:"required"`
}

type UpdatePaymentRequest struct {
	PaymentStatus string `json:"payment_status" binding:"required"`
}

