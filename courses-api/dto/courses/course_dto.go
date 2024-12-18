package courses

//dto es domain

type CoursesResponse_Full []CourseResponse_Full // si está en plural son muchos

type CoursesResponse_Result []CourseResponse_Result

type CoursesRequest_Category []CourseRequest_Category

type CoursesRequest_Title []CourseRequest_Title

type CoursesRequest_Description []CourseRequest_Description

type CoursesNewRequest []CourseNewRequest

type CoursesNewResponse []CourseNewResponse

// BUSQUEDA DE CURSOS

type CourseRequest_Category struct {
	Category string `json:"category"`
}

type CourseRequest_Title struct {
	Title string `json:"title"`
}

type CourseRequest_Description struct {
	Description string `json:"description"`
}

type CourseRequest_Token struct {
	Token string `json:"token"`
}

// RESULTADO LISTANDO CURSOS

type CourseResponse_Result struct {
	ID_Course   string  `json:"id_course"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ImageURL    string  `json:"image_url"`
	Rating      float64 `json:"rating"`
}

// DETALLE DEL CURSO

type CourseResponse_Full struct {
	ID_Course    string `json:"id_course"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Category     string `json:"category"`
	ImageURL     string `json:"image_url"`
	Duration     string `json:"duration"`
	Instructor   string `json:"instructor"`
	Points       string `json:"points"`
	Capacity     int    `json:"capacity"`
	Requirements string `json:"requirements"`
}

// // CREAR NUEVO CURSO

type CourseNewRequest struct {
    Title        string `json:"title"`
    Description  string `json:"description"`
    Category     string `json:"category"`
    ImageURL     string `json:"image_url"`
    Duration     string `json:"duration"`
    Instructor   string `json:"instructor"`
    Points       int    `json:"points"`
    Capacity     int    `json:"capacity"`
    Requirements string `json:"requirements"`
}

type CourseNewResponse struct {
	Token string `json:"token"`
}

// INSCRIPCION EN CURSO
type CourseRequest_Registration struct {
	Token     string `json:"token"`
	ID_Course string `json:"id_course"`
}

type CourseResponse_Registration struct {
	ID_Course string `json:"id_course"`
}

type CourseNew struct {
	ID_Course string `json:"id_course"`
}