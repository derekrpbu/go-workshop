package models

type User struct {
	UserID   int64  `json:"user_id,string,omitempty" bson:"user_id,string,omitempty"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	IsActive bool   `json:"is_active,string,omitempty" bson:"is_active,string,omitempty"`
}
