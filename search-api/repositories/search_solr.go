package solrClient

import (
    "github.com/rtt/Go-Solr"
    "log"
)

var SolrClient *solr.Connection

func init() {
    var err error
    SolrClient, err = solr.Init("http://localhost:8983/solr/courses")
    if err != nil {
        log.Fatalf("Failed to initialize Solr client: %v", err)
    }
}

// UpdateDocument actualiza un documento en SolR
func UpdateDocument(document map[string]interface{}) (*solr.UpdateResponse, error) {
    resp, err := SolrClient.Update(document, true)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

// SelectRaw ejecuta una consulta sin procesar en SolR
func SelectRaw(query string) (*solr.SelectResponse, error) {
    resp, err := SolrClient.SelectRaw(query)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

// Select ejecuta una consulta estructurada en SolR
func Select(query *solr.Query) (*solr.SelectResponse, error) {
    resp, err := SolrClient.Select(query)
    if err != nil {
        return nil, err
    }
    return resp, nil
}