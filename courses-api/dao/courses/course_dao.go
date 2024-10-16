package courses

type Course struct {
	ID_Course    string `bson:"id_course"`
	Title        string `bson:"title"`
	Description  string `bson:"description"`
	Category     string `bson:"category"`
	ImageURL     string `bson:"image_url"`
	Duration     string `bson:"duration"`
	Instructor   string `bson:"instructor"`
	Points       string `bson:"points"`
	Capacity     int    `bson:"capacity"`
	Requirements string `bson:"requirements"`
}
