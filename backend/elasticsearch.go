package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/olivere/elastic/v7"
)

const (
	ES_URL = "https://127.0.0.1:9200/" // demo url
)

func readFromES(query elastic.Query, index string) (*elastic.SearchResult, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(ES_URL),
		elastic.SetBasicAuth("elastic", "hZrWYXZ-BU01PcaTrWae"), // fake password
		elastic.SetHealthcheck(false),
		elastic.SetSniff(false),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "ELASTIC ", log.LstdFlags)),
		elastic.SetTraceLog(log.New(os.Stdout, "ELASTIC ", log.LstdFlags)))
	if err != nil {
		return nil, err
	}

	searchResult, err := client.Search().
		Index(index).
		Query(query).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return nil, err
	}
	fmt.Println("read check")
	return searchResult, nil
}

func saveToES(i interface{}, index string, id string) error {
	client, err := elastic.NewClient(
		elastic.SetURL(ES_URL),
		elastic.SetBasicAuth("elastic", "hZrWYXZ-BU01PcaTrWae"),
		elastic.SetHealthcheck(false),
		elastic.SetSniff(false),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "ELASTIC ", log.LstdFlags)),
		elastic.SetTraceLog(log.New(os.Stdout, "ELASTIC ", log.LstdFlags))) // fake password
	if err != nil {
		return err
	}

	_, err = client.Index().
		Index(index).
		Id(id).
		BodyJson(i).
		Do(context.Background())
	return err
}

func deleteFromES(query elastic.Query, index string) error {
	client, err := elastic.NewClient(
		elastic.SetURL(ES_URL),
		elastic.SetBasicAuth("elastic", "hZrWYXZ-BU01PcaTrWae"),
		elastic.SetHealthcheck(false),
		elastic.SetSniff(false),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "ELASTIC ", log.LstdFlags)),
		elastic.SetTraceLog(log.New(os.Stdout, "ELASTIC ", log.LstdFlags))) // fake password
	if err != nil {
		return err
	}

	_, err = client.DeleteByQuery().
		Index(index).
		Query(query).
		Pretty(true).
		Do(context.Background())

	return err
}
