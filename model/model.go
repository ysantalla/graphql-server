package model

// User type
type User struct {
	ID    string  `json:"id,omitempty" bson:"_id"`
	Name  string  `json:"name,omitempty" bson:"name"`
	Email string  `json:"email,omitempty" bson:"name"`
	Role  []*Role `json:"role,omitempty"`
}

//Role type
type Role struct {
	ID   string  `json:"id,omitempty" bson:"_id"`
	Name string  `json:"name,omitempty" bson:"name"`
	User []*User `json:"user,omitempty"`
}

//File type
type File struct {
	ID       string `json:"id,omitempty" bson:"_id"`
	Path     string `json:"path,omitempty" bson:"path"`
	Filename string `json:"filename,omitempty" bson:"filename"`
	Mimetype string `json:"mimetype,omitempty" bson:"mimetype"`
	Encoding string `json:"encoding,omitempty" bson:"encoding"`
	Size     int    `json:"size,omitempty" bson:"size"`
}
