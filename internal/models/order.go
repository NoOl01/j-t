package models

type Order struct {
	Id         int64       `gorm:"primaryKey" json:"id"`
	UserId     int64       `gorm:"not null;index" json:"user_id"`
	User       User        `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"order_items"`
	Status     string      `gorm:"not null" json:"status"`
	Price      float64     `gorm:"not null" json:"price"`
}
