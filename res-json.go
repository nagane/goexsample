package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type BacklogAddIssue struct {
	created    string         `json:"created"`
	project    BacklogProject `json:"project"`
	id         int
	issue_type string `json:"type"`
	content    BacklogContent
	//resolution resolution
	//actualHours int
	//issueType issueType
	//milestone mileStone
	//versions versions
	notifications BacklogNotifications
	createdUser   BacklogCreatedUser
}

type BacklogProject struct {
	archived          bool
	projectKey        string
	name              string
	chartEnabled      bool
	id                int
	subtaskingEnabled bool
}

type BacklogContent struct {
	summary string
	keyId   int `json:"key_id"`
	customFields
}

type BacklogNotifications struct {
}

type BacklogCreatedUser struct {
}

func main() {

	bytes, err := ioutil.ReadFile("temp/backlog.json")
	//bytes, err := ioutil.ReadFile("temp/test.json")
	if err != nil {
		log.Fatal(err)
	}

	var issues []BacklogAddIssue
	if err := json.Unmarshal(bytes, &issues); err != nil {
		log.Fatal(err)
	}

	for _, p := range issues {
		fmt.Printf("%s", p.created)
	}
}
