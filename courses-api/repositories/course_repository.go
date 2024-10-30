package repositories

import (
    "context"
    "courses-api/dao"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type CourseRepository struct {
    collection *mongo.Collection
}

func NewCourseRepository(client *mongo.Client, dbName, collectionName string) *CourseRepository {
    collection := client.Database(dbName).Collection(collectionName)
    return &CourseRepository{collection: collection}
}

func (r *CourseRepository) CreateCourse(course dao.Course) error {
    _, err := r.collection.InsertOne(context.Background(), course)
    return err
}

func (r *CourseRepository) UpdateCourse(id string, course dao.Course) error {
    filter := bson.M{"id_course": id}
    update := bson.M{"$set": course}
    _, err := r.collection.UpdateOne(context.Background(), filter, update)
    return err
}

func (r *CourseRepository) DeleteCourse(id string) error {
    filter := bson.M{"id_course": id}
    _, err := r.collection.DeleteOne(context.Background(), filter)
    return err
}

func (r *CourseRepository) GetCourse(id string) (dao.Course, error) {
    var course dao.Course
    filter := bson.M{"id_course": id}
    err := r.collection.FindOne(context.Background(), filter).Decode(&course)
    return course, err
}