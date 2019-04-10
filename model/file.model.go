package model

//File type
type File struct {
	ID       string `json:"id,omitempty" bson:"_id"`
	Path     string `json:"path,omitempty" bson:"path"`
	Filename string `json:"filename,omitempty" bson:"filename"`
	Mimetype string `json:"mimetype,omitempty" bson:"mimetype"`
	Encoding string `json:"encoding,omitempty" bson:"encoding"`
	Size     int    `json:"size,omitempty" bson:"size"`
}
