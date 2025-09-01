package models

type Points struct {
	Id     int64 `gorm:"primaryKey"`
	Value  int   `gorm:"not null"`
	UserId int64 `gorm:"not null;uniqueIndex"`
	User   User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
