package model

import (
	"time"
)

type User struct {
	User_Id      string    `gorm:"size:64;primary_key"`
	User_Name    string    `gorm:"type:varchar(20);not null"`
	User_Pwd     string    `gorm:"size:255;not null"`
	User_Email   string    `gorm:"type:varchar(32);not null"`
	User_Tel     string    `gorm:"type:varchar(11)"`
	GithubId     string    `gorm:"size:64;uniqueIndex:github"`
	User_Logo    string    `gorm:"type:varchar(255)"`
	Created_date time.Time `gorm:"DEFAULT:CURRENT_TIMESTAMP"`
	Updated_date time.Time `gorm:"DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
