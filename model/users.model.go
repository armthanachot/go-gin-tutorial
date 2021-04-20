package model
// model
type User struct {
	Id        uint   `gorm:"primary_key" json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       string    `json:"age"`
	Email     string `json:"email"`
}
