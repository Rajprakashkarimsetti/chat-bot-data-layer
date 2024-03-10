package models

type QueryInfo struct {
	Id       string `json:"id,omitempty"`
	Question string `json:"question,omitempty"`
	Solution string `json:"solution,omitempty"`
	Count    int64  `json:"count,omitempty"`
}

type UserInfo struct {
	ID        string `gorm:"unique" json:"id,omitempty`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Dob       string `json:"dob,omitempty"`
	Phone     string `gorm:"unique" json:"phone,omitempty"`
	Gender    string `json:"gender,omitempty"`
	Email     string `gorm:"unique" json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
}
