package app

import (
	courseController "courses-api/controllers/courses"

	log "github.com/sirupsen/logrus"
)

// MapUrls asigna las rutas a los controladores correspondientes
func MapUrls() {

	// Courses Mapping
	router.GET("/course", courseController.GetAll)
	router.GET("/course/:id", courseController.GetCourseByID)
	router.GET("/course/title=:title", courseController.GetCourseByTitle)
	router.GET("/course/category=:category", courseController.GetCourseByCategory)
	router.GET("/course/description=:description", courseController.GetCourseByDescription)

	//router.POST("/course/register", courseController.RegisterUserToCourse)
	router.POST("/course", courseController.CreateCourse)
	router.POST("/course", courseController.PostCourse)

	//router.PUT("/course/:id_course", courseController.PutCourse)
	router.PUT("/course/:id_course", courseController.Update) // no debería solo dejar este put? Igual dentro de Put se llama a update

	router.DELETE("/course", courseController.Delete)

	log.Info("Finishing mappings configurations")
}
