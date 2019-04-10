package model

//Role type
type Role struct {
	ID   string  `json:"id,omitempty" bson:"_id"`
	Name string  `json:"name,omitempty" bson:"name"`
	User []*User `json:"user,omitempty"`
}
