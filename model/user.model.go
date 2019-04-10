package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User type
type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name  string             `json:"name,omitempty"`
	Email string             `json:"email,omitempty"`
	Role  []*Role            `json:"role,omitempty"`
}
