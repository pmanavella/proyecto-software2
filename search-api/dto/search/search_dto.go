package search_dto

type SearchDto struct {
	ID_Course    string    `json:"id_course"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Category     string `json:"category"`
	ImageURL     string `json:"image_url"`
	Duration     string `json:"duration"`
	Requirements string `json:"requirements"`
	Instructor  string `json:"instructor"`
}

