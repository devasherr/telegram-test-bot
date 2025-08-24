package models

type User struct {
	ID          int64  `bson:"_id"`
	FirstName   string `bson:"first_name"`
	LastName    string `bson:"last_name"`
	UserName    string `bson:"username"`
	PhoneNumber string `bson:"phone_number"`
	Registerd   bool   `bson:"registered"`
}
