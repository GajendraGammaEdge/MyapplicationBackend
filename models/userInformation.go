package models

type UserInformation struct {
	Id       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Password string `gorm:"type:varchar(255)" json:"-"`
	Address  string `json:"address"`
}

func (UserInformation) TableName() string {
	return "user_information"
}
