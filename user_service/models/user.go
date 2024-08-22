package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents the structure of a user in your application
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
}
