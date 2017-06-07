package main

import (
	"fmt"
	"github.com/antonholmquist/jason"
	"io/ioutil"
	"log"
)

func main() {

	bytes, err := ioutil.ReadFile("temp/backlog.json")
	//bytes, err := ioutil.ReadFile("temp/test.json")
	if err != nil {
		log.Fatal(err)
	}

	v, _ := jason.NewObjectFromBytes(bytes)

	summary, _ := v.GetString("content", "summary")
	// person in charge
	picName, _ := v.GetString("content", "assignee", "name")
	dueDate, _ := v.GetString("content", "dueDate")
	keyNum, _ := v.GetString("project", "projectKey")

	respons_str := `
    追加されたタスク：` + summary +
		`
    担当者：` + picName +
		`
    期日：` + dueDate +
		`
    チケット番号：` + keyNum

	fmt.Printf("%s", respons_str)

}
