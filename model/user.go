package model

type User struct {
	UserId         string     `json:"UserId"`
	IdentityNumber string     `json:"IdentityNumber"`
	Username       string     `json:"Username"`
	DateOfBirth    string     `json:"DateOfBirth"`
	Profession     Profession `json:"Profession"`
	Education      Education  `json:"Education"`
	UserStatus     int        `json:"UserStatus"`
}

type Profession struct {
	ProfessionId   string `json:"ProfessionId"`
	ProfessionName string `json:"ProfessionName"`
}

type Education struct {
	EducationId   string `json:"EducationId"`
	EducationName string `json:"EducationName"`
}
