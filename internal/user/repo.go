package user

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repo struct {
	col *mongo.Collection
}

func (r *Repo) FindByEmail(ctx context.Context, email string) (*User, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	filter := bson.M{"email": email}
	var user User
	err := r.col.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &User{}, fmt.Errorf("user not found")
		}
		return &User{}, fmt.Errorf("failed to find user: %w", err)
	}
	return &user, nil

}

func NewRepo(db *mongo.Database) *Repo {
	return &Repo{col: db.Collection("users")}
}

func (r *Repo) Create(ctx context.Context, u *User) (*User, error) {
	// 1. Manually generate the ObjectID before inserting
	if u.ID.IsZero() {
		u.ID = bson.NewObjectID()
	}

	// 2. Insert the user (who now already has an ID)
	_, err := r.col.InsertOne(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// 3. Since we set the ID ourselves, we don't need to check res.InsertedID
	return u, nil
}
