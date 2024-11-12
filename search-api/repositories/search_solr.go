package repositories

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "github.com/stevenferrer/solr-go"
    courses "search-api/dao"
)

type SolrConfig struct {
    Host       string // Solr host
    Port       string // Solr port
    Collection string // Solr collection name
}

type Solr struct {
    Client     *solr.JSONClient
    Collection string
}

// NewSolr initializes a new Solr client
func NewSolr(config SolrConfig) Solr {
    // Construct the BaseURL using the provided host and port
    baseURL := fmt.Sprintf("http://%s:%s", config.Host, config.Port)
    client := solr.NewJSONClient(baseURL)

    return Solr{
        Client:     client,
        Collection: config.Collection,
    }
}

// Index adds a new course document to the Solr collection
func (searchEngine Solr) Index(ctx context.Context, course courses.Search) (string, error) {
    // Prepare the document for Solr
    doc := map[string]interface{}{
        "id":           course.ID_Course,
        "description":  course.Description,
        "capacity":     course.Capacity,
        "instructor":   course.Instructor,
        "category":     course.Category,
        "duration":     course.Duration,
    }   

    // Prepare the index request
    indexRequest := map[string]interface{}{
        "add": []interface{}{doc}, // Use "add" with a list of documents
    }

    // Index the document in Solr
    body, err := json.Marshal(indexRequest)
    if err != nil {
        return "", fmt.Errorf("error marshaling course document: %w", err)
    }

    // Index the document in Solr
    resp, err := searchEngine.Client.Update(ctx, searchEngine.Collection, solr.JSON, bytes.NewReader(body))
    if err != nil {
        return "", fmt.Errorf("error indexing course: %w", err)
    }
    if resp.Error != nil {
        return "", fmt.Errorf("failed to index course: %v", resp.Error)
    }

    // Commit the changes
    if err := searchEngine.Client.Commit(ctx, searchEngine.Collection); err != nil {
        return "", fmt.Errorf("error committing changes to Solr: %w", err)
    }

    return course.ID_Course, nil
}

// Update modifies an existing course document in the Solr collection
func (searchEngine Solr) Update(ctx context.Context, course courses.Search) error {
    // Prepare the document for Solr
    doc := map[string]interface{}{
        "id":         course.ID_Course,
        "description":       course.Description,
        "capacity":   course.Capacity,
        "instructor": course.Instructor,
        "category":     course.Category,
        "duration":     course.Duration,
    }

    // Prepare the update request
    updateRequest := map[string]interface{}{
        "add": []interface{}{doc}, // Use "add" with a list of documents
    }

    // Update the document in Solr
    body, err := json.Marshal(updateRequest)
    if err != nil {
        return fmt.Errorf("error marshaling course document: %w", err)
    }

    // Execute the update request using the Update method
    resp, err := searchEngine.Client.Update(ctx, searchEngine.Collection, solr.JSON, bytes.NewReader(body))
    if err != nil {
        return fmt.Errorf("error updating course: %w", err)
    }
    if resp.Error != nil {
        return fmt.Errorf("failed to update course: %v", resp.Error)
    }

    // Commit the changes
    if err := searchEngine.Client.Commit(ctx, searchEngine.Collection); err != nil {
        return fmt.Errorf("error committing changes to Solr: %w", err)
    }

    return nil
}

// Delete removes a course document from the Solr collection
func (searchEngine Solr) Delete(ctx context.Context, id string) error {
    // Prepare the delete request
    docToDelete := map[string]interface{}{
        "delete": map[string]interface{}{
            "id": id,
        },
    }

    // Update the document in Solr
    body, err := json.Marshal(docToDelete)
    if err != nil {
        return fmt.Errorf("error marshaling course document: %w", err)
    }

    // Execute the delete request using the Update method
    resp, err := searchEngine.Client.Update(ctx, searchEngine.Collection, solr.JSON, bytes.NewReader(body))
    if err != nil {
        return fmt.Errorf("error deleting course: %w", err)
    }
    if resp.Error != nil {
        return fmt.Errorf("failed to delete course: %v", resp.Error)
    }

    // Commit the changes
    if err := searchEngine.Client.Commit(ctx, searchEngine.Collection); err != nil {
        return fmt.Errorf("error committing changes to Solr: %w", err)
    }

    return nil
}

// Search executes a search query in the Solr collection
func (searchEngine Solr) Search(ctx context.Context, query string, limit int, offset int) ([]courses.Search, error) {
    // Prepare the Solr query with limit and offset
    solrQuery := fmt.Sprintf("q=(name:%s)&rows=%d&start=%d", query, limit, offset)

    // Execute the search request
    resp, err := searchEngine.Client.Query(ctx, searchEngine.Collection, solr.NewQuery(solrQuery))
    if err != nil {
        return nil, fmt.Errorf("error executing search query: %w", err)
    }
    if resp.Error != nil {
        return nil, fmt.Errorf("failed to execute search query: %v", resp.Error)
    }

    // Parse the response and extract course documents
    var coursesList []courses.Search
    for _, doc := range resp.Response.Documents {
        course := courses.Search{
            ID_Course:         getStringField(doc, "id"),
            Description:       getStringField(doc, "description"),
            Capacity:   getIntField(doc, "capacity"),
            Instructor: getStringField(doc, "instructor"),
            Category:     getStringField(doc, "category"),
            Duration:     getIntField(doc, "duration"),
        }
        coursesList = append(coursesList, course)
    }

    return coursesList, nil
}

// Helper function to safely get string fields from the document
func getStringField(doc map[string]interface{}, field string) string {
    if val, ok := doc[field].(string); ok {
        return val
    }
    if val, ok := doc[field].([]interface{}); ok && len(val) > 0 {
        if strVal, ok := val[0].(string); ok {
            return strVal
        }
    }
    return ""
}

// Helper function to safely get int fields from the document
func getIntField(doc map[string]interface{}, field string) int {
    if val, ok := doc[field].(float64); ok {
        return int(val)
    }
    if val, ok := doc[field].([]interface{}); ok && len(val) > 0 {
        if intVal, ok := val[0].(float64); ok {
            return int(intVal)
        }
    }
    return 0
}