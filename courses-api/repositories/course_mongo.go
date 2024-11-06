package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	dao "courses-api/dao/courses"
)

type MongoConfig struct {
	Host       string
	Port       string
	Username   string
	Password   string
	Database   string
	Collection string
}

type Mongo struct {
	client     *mongo.Client
	database   string
	collection string
}

const (
	connectionURI = "mongodb://%s:%s"
)

func NewMongo(config MongoConfig) Mongo {
	credentials := options.Credential{
		Username: config.Username,
		Password: config.Password,
	}

	ctx := context.Background()
	uri := fmt.Sprintf(connectionURI, config.Host, config.Port)
	cfg := options.Client().ApplyURI(uri).SetAuth(credentials)

	client, err := mongo.Connect(ctx, cfg)
	if err != nil {
		log.Panicf("error connecting to mongo DB: %v", err)
	}

	return Mongo{
		client:     client,
		database:   config.Database,
		collection: config.Collection,
	}
}

func (repository Mongo) GetCourseByID(ctx context.Context, id string) (dao.Course, error) {
	// Get from MongoDB
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dao.Course{}, fmt.Errorf("error converting id to mongo ID: %w", err)
	}
	result := repository.client.Database(repository.database).Collection(repository.collection).FindOne(ctx, bson.M{"_id": objectID})
	if result.Err() != nil {
		return dao.Course{}, fmt.Errorf("error finding document: %w", result.Err())
	}

	// Convert document to DAO
	var cursoDAO dao.Course
	if err := result.Decode(&cursoDAO); err != nil {
		return dao.Course{}, fmt.Errorf("error decoding result: %w", err)
	}
	return cursoDAO, nil
}

func (repository Mongo) Create(ctx context.Context, curso dao.Course) (string, error) {
	result, err := repository.client.Database(repository.database).Collection(repository.collection).InsertOne(ctx, curso)
	if err != nil {
		return "", fmt.Errorf("error creating document: %w", err)
	}
	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("error converting mongo ID to object ID")
	}
	return objectID.Hex(), nil
}

func (repository Mongo) Update(ctx context.Context, course dao.Course) error {
	// Convert curso ID to MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(course.ID_Course)
	if err != nil {
		return fmt.Errorf("error converting id to mongo ID: %w", err)
	}

	// Create an update document
	update := bson.M{}

	// Only set the fields that are not empty or their default value
	if course.Description != "" {
		update["Descripcion"] = course.Description
	}
	if course.Title != "" {
		update["Título"] = course.Title
	}
	if course.Category != "" {
		update["Categoria"] = course.Category
	}
	if course.Requirements != "" {
		update["Requisitos"] = course.Requirements
	}
	if course.Points != "" { 
		update["Puntos"] = course.Points
	}

	// Update the document in MongoDB
	if len(update) == 0 {
		return fmt.Errorf("no fields to update for curso ID %s", course.ID_Course)
	}

	filter := bson.M{"_id": objectID}
	result, err := repository.client.Database(repository.database).Collection(repository.collection).UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return fmt.Errorf("error updating document: %w", err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("no document found with ID %s", course.ID_Course)
	}

	return nil
}