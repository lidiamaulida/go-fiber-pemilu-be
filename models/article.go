package models

type Article struct { 
	Id          int64  `goerm:"primaryKey:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(250)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Date        string `gorm:"type:date" json:"date"`
	Image       string `gorm:"type:varchar(255)" json:"image"`
	UserID      int64  `json:"userId"`
	User        UserResponse `gorm:"foreignKey:UserID" json:"user"`
}

type ArticleResponse struct {
	ID          int    `json:"id"`
	ArticleName string `json:"articleName"`
	Description string `json:"description"`
	Image       string `json:"image"`
	UserID      int    `json:"userId"`
}

func (ArticleResponse) TableName() string {
	return "articles"
}
