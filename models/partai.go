package models

type Partai struct {
	Id                int64   `goerm:"primaryKey:auto_increment" json:"id"`
	Name              string  `gorm:"type:varchar(250)" json:"name"`
	Chairman          string  `gorm:"type:varchar(250)" json:"chairman"`
	VisionAndMission  string  `gorm:"type:varchar(250)" json:"visionAndMission"`
	Address           string  `goerm:"type:varchar(250)" json:"address"`
	Image             string  `gorm:"type:varchar(255)" json:"image"`
}