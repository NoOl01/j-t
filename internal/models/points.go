package models

type Points struct {
	Id     int64 `gorm:"primaryKey" json:"-"`
	Value  int   `gorm:"not null" json:"value"`
	UserId int64 `gorm:"not null;uniqueIndex" json:"-"`
}
