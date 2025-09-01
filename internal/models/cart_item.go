package models

type CartItem struct {
	Id        int64   `gorm:"primaryKey" json:"id"`
	ProductId int64   `gorm:"not null;index" json:"product_id"`
	Product   Product `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:CASCADE" json:"product"`
	UserId    int64   `gorm:"not null;index" json:"user_id"`
	User      User    `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:CASCADE" json:"user"`
	Count     int     `gorm:"not null"`
}
