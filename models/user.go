package models

type User struct{
	Id         int64   `goerm:"primaryKey:auto_increment" json:"id"`
	Fullname   string  `goerm:"type:varchar(250)" json:"fullname"`
	Address    string  `goerm:"type:varchar(250)" json:"address"`
	Gender     string  `goerm:"type:varchar(250)" json:"gender"`
	Username   string  `goerm:"type:varchar(250)" json:"username"`
	Password   string  `goerm:"type:varchar(250)" json:"password"`
	Role       string  `goerm:"type:varchar(250)" json:"role"`
}