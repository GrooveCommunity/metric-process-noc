package main

import (
	"encoding/json"
	storage "github.com/GrooveCommunity/glib-cloud-storage/gcp"
	"github.com/GrooveCommunity/glib-noc-event-structs/gcs"
	"log"
)

func main() {
	var request gcs.IssuesMetric

	payload := storage.GetObject("noc-paygo","jira-issues.json")

	err := json.Unmarshal(payload, &request)

	log.Println(err)

	//log.Println(request)
}