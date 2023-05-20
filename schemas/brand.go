package schemas

import "go.mongodb.org/mongo-driver/bson/primitive"

type Brand struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name,omitempty"`
	Code           string             `bson:"code,omitempty"`
	ImagePath      string             `bson:"image_path,omitempty"`
	ImageExtension string             `bson:"image_extension,omitempty"`
	ImageIndex     int                `bson:"image_index"`
	Incomplete     bool               `bson:"incomplete"`
	RetryImage     bool               `bson:"retry_image"`
}
