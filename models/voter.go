package models

type Voter struct {
    ID       int            `json:"id" gorm:"primaryKey:autoIncrement"`
    UserID   int            `json:"userId"`
    PaslonID int            `json:"paslonId" gorm:"uniqueIndex:idx_user_paslon"`
    User     UserResponse   `json:"user" gorm:"foreignKey:UserID"`  
    Paslon   PaslonVoterResponse         `json:"paslon" gorm:"foreignKey:PaslonID"`  
}

type VoterResponse struct {
    ID       int `json:"id"`  
    Paslon   PaslonVoterResponse
}
