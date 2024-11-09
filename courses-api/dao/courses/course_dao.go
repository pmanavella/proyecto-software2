package courses

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	ID_Course    primitive.ObjectID `bson:"_id,omitempty"`
	Title        string             `bson:"title"`
	Description  string             `bson:"description"`
	Category     string             `bson:"category"`
	ImageURL     string             `bson:"image_url"`
	Duration     string             `bson:"duration"`
	Instructor   string             `bson:"instructor"`
	Points       string             `bson:"points"`
	Capacity     int                `bson:"capacity"`
	Requirements string             `bson:"requirements"`
}
