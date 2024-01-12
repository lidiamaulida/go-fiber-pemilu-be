package models

type Article struct { 
	Id          int64  `goerm:"primaryKey:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(250)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Author      string `gorm:"type:varchar(250)" json:"author"`
	Date        string `gorm:"type:date" json:"date"`
	Image       string `gorm:"type:varchar(255)" json:"image"`
}