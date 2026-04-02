package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ID primitive.objectID `bson:"_id,omitempty" json:"id"`

	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"-"`

	Role string `bson:"role" json:"role"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type PublicUser struct {
	ID        string    `json:"id"`
	Role      string    `json:"role"`
	Role      string    ` json:"role"`
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