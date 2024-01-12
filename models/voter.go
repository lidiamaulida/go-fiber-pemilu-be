package models 

type Voter struct {
	Id         int64   `goerm:"primaryKey:auto_increment" json:"id"`
	Fullname   string  `goerm:"type:varchar(250)" json:"fullname"`
	Address    string  `goerm:"type:varchar(250)" json:"address"`
	Gender     string  `goerm:"type:varchar(250)" json:"gender"`
	Paslon     string  `goerm:"type:varchar(250)" json:"paslon"`
}