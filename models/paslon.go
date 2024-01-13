package models

type Paslon struct {
	Id                int64   `goerm:"primaryKey:auto_increment" json:"id"`
	Name              string  `gorm:"type:varchar(250)" json:"name"`
	No                int64   `gorm:"type:int" json:"no"`
	VisionAndMission  string  `gorm:"type:varchar(250)" json:"visionAndMission"`
	Image             string  `gorm:"type:varchar(250)" json:"image"`
	Partais         []PartaiResponse `json:"partais" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type PaslonVoterResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	OrderNum int    `json:"orderNum"`
}

func (PaslonVoterResponse) TableName() string {
	return "paslons"
}