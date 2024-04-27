package models

type User struct {
	ID   string `json:"_id" bson:"_id"`
	UserId string `json:"user_id" bson:"user_id"`
	Email string `json:"email" bson:"email"`
	Name string `json:"name" bson:"name"`
}
