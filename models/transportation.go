package models

type Transportation struct {
	ID                   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	DestinationID        int    `gorm:"not null;index" json:"destination_id"`
	TransportTypeID      int    `gorm:"not null;index" json:"transport_type_id"`
	Price                int    `gorm:"not null" json:"price"`
	Estimate             string `gorm:"type:varchar(255)" json:"estimate"`
	DetailTransportation string `gorm:"column:detail_tranportation;type:varchar(255)" json:"detail_transportation"`
}

func (Transportation) TableName() string {
	return "transportation"
}

type CreateTransportationRequest struct {
	DestinationID        int    `json:"destination_id" binding:"required"`
	TransportTypeID      int    `json:"transport_type_id" binding:"required"`
	Price                int    `json:"price" binding:"required"`
	Estimate             string `json:"estimate"`
	DetailTransportation string `json:"detail_transportation"`
}

type UpdateTransportationRequest struct {
	Price                *int    `json:"price"`
	Estimate             *string `json:"estimate"`
	DetailTransportation *string `json:"detail_transportation"`
}

type AccommodationResponse struct {
	DestinationID        int    `json:"destination_id"`
	DestinationName      string `json:"destination_name"`
	DestinationImage     string `json:"destination_image"`
	TransportID          int    `json:"transport_id"`
	TransportTypeName    string `json:"transport_type_name"`
	Price                int    `json:"price"`
	Estimate             string `json:"estimate"`
	DetailTransportation string `json:"detail_transportation"`
}

type TransportType struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}

func (TransportType) TableName() string {
	return "transport_types"
}

type CreateTransportTypeRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTransportTypeRequest struct {
	Name string `json:"name" binding:"required"`
}
