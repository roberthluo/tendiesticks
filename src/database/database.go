package database

import (
	"context"
	// "encoding/json"
	"fmt"
	// "reflect"
	// "time"

	"github.com/olivere/elastic"
)

// Tweet is a structure used for serializing/deserializing data in Elasticsearch.
type Tweet struct {
	Metadata 	string                `json:"user"`
	TimeSeries  string                `json:"message"`
	// Retweets int                   `json:"retweets"`
	// Image    string                `json:"image,omitempty"`
	// Created  time.Time             `json:"created,omitempty"`
	// Tags     []string              `json:"tags,omitempty"`
	// Location string                `json:"location,omitempty"`
	// Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}



const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"tweet":{
			"properties":{
				"user":{
					"type":"keyword"
				},
				"message":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"image":{
					"type":"keyword"
				},
				"created":{
					"type":"date"
				},
				"tags":{
					"type":"keyword"
				},
				"location":{
					"type":"geo_point"
				},
				"suggest_field":{
					"type":"completion"
				}
			}
		}
	}
}`


func SetupDB(){
	// Starting with elastic.v5, you must pass a context to execute each service
	ctx := context.Background()

	// Obtain a client and connect to the default Elasticsearch installation
	// on 127.0.0.1:9200. Of course you can configure your client to connect
	// to other hosts and configure it in various other ways.
	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}