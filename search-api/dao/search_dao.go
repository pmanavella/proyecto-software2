package dao_search


type Search struct {
	ID_Course    string    `bson:"id_course"`
	Description  string    `bson:"description"`
	Category     string    `bson:"category"`
	ImageURL     string    `bson:"image_url"`
	Duration     string    `bson:"duration"`
	Instructor   string    `bson:"instructor"`
	Points       string    `bson:"points"`
	Requirements string    `bson:"requirements"`
	Capacity 	 int 	   `bson:"capacity"`
}

type CourseNew struct {
	Operation string `json:"operation"`
	CourseID   string `json:"hotel_id"`
}