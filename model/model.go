package model

type User struct {
	ID int64 `gorm:"primaryKey"`
}

type Profile struct {
	ID     int64  `gorm:"primaryKey"`
	Name0  string `gorm:"column:name_0;not null;uniqueIndex:name_idx"`
	Name1  string `gorm:"column:name_1;not null;unique:name_idx"`
	Name2  string `gorm:"column:name_2;not null;idx:name_idx"`
	Name3  string `gorm:"column:name_3;not null;idx:name_idx;unique"`
	Name4  string `gorm:"column:name_4;not null;index:name_idx,unique"`
	Name5  string `gorm:"column:name_5;not null;index:name_idx;unique"`
	Name6  string `gorm:"column:name_6;not null;index:name_idx;uniqueIndex"`
	Name7  string `gorm:"column:name_7;not null;uniqueIndex"`
	UserID int64  `gorm:"column:user_id"`
	User   User   `gorm:"foreignKey:UserID"`
}
