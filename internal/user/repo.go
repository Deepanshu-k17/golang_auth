package user

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Repo struct {
	col *mongo.Collection
}

func (r *Repo) FindByEmail(ctx context.Context, email string) (User, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	filter := bson.M{"email": email}
	var user User
	err := r.col.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return User{}, fmt.Errorf("user not found")
		}
		return User{}, fmt.Errorf("failed to find user: %w", err)
	}
	return user, nil

}

func NewRepo(db *mongo.Database) *Repo {
	return &Repo{col: db.Collection("users")}
}

func (r *Repo) Create(ctx context.Context, u *User) (User, error) {
	res, err := r.col.InsertOne(ctx, u)
	if err != nil {
		return User{}, fmt.Errorf("failed to create user: %w")
	}
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return User{}, fmt.Errorf("Insert user failed and inserted ID is not an ObjectID")
	}
	u.ID = id
	return u, nil
}
