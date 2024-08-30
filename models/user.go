package models

import (
	DTO "example/main/dto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

func (u *User) ToDTO() DTO.UserProfileDTO {
	return DTO.UserProfileDTO{
		ID:    u.ID.Hex(), // Convert ObjectID to string
		Name:  u.Name,
		Email: u.Email,
	}
}
