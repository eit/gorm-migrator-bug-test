package model

type User struct {
	ID int64 `gorm:"primaryKey"`
}

type Profile struct {
	ID   int64 `gorm:"primaryKey"`
	User User  `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// UserID int64 `gorm:"column:user_id"`
	// User   User  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
