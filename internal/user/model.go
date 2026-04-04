package user

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Email     string        `bson:"email" json:"email"`
	Password  string        `bson:"password" json:"-"`
	Role      string        `bson:"role" json:"role"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
}

type PublicUser struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"` // Added this so ToPublic works
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToPublic(u User) PublicUser {
	return PublicUser{
		ID:        u.ID.Hex(),
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
