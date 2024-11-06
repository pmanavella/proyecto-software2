package client

import (
    solrClient "Search/solr"
    "errors"
    "github.com/rtt/Go-Solr"
    log "github.com/sirupsen/logrus"
    "strconv"
)

type courseClient struct{}

type courseClientInterface interface {
    UpdateCourse(document map[string]interface{}) error
    GetCourses() (*solr.DocumentCollection, error)
    GetCourseById(id string) (*solr.DocumentCollection, error)
    GetCoursesByCapacity(minCapacity, maxCapacity int) (*solr.DocumentCollection, error)
}

var SolrCourseClient courseClientInterface

func init() {
    SolrCourseClient = &courseClient{}
}

func (c courseClient) UpdateCourse(document map[string]interface{}) error {
    resp, err := solrClient.SolrClient.Update(document, true)
    if err != nil {
        return err
    }
    log.Printf("Solr Response: %s", resp.String())

    return nil
}

func (c courseClient) GetCourses() (*solr.DocumentCollection, error) {
    q := "q=*:*"

    resp, err := solrClient.SolrClient.SelectRaw(q)
    if err != nil {
        log.Info(err)
        return &solr.DocumentCollection{}, err
    }

    result := resp.Results
    return result, nil
}

func (c courseClient) GetCourseById(id string) (*solr.DocumentCollection, error) {
    q := solr.Query{
        Params: solr.URLParamMap{
            "q": []string{"id:" + id},
        },
    }

    resp, err := solrClient.SolrClient.Select(&q)
    if err != nil {
        return &solr.DocumentCollection{}, err
    }

    result := resp.Results
    if result.Len() == 0 {
        return &solr.DocumentCollection{}, errors.New("course not found")
    }

    return result, nil
}

func (c courseClient) GetCoursesByCapacity(minCapacity, maxCapacity int) (*solr.DocumentCollection, error) {
    q := solr.Query{
        Params: solr.URLParamMap{
            "q": []string{"capacity:[" + strconv.Itoa(minCapacity) + " TO " + strconv.Itoa(maxCapacity) + "]"},
        },
    }

    resp, err := solrClient.SolrClient.Select(&q)
    if err != nil {
        return &solr.DocumentCollection{}, err
    }

    result := resp.Results
    return result, nil
}