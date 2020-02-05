package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"log"
)


func main() {

	// Generate the submission in a pretty JSON format
	requestBody, err := json.Marshal(map[string]string{
		"name" : "Karl Mathias Moberg",
		"email" : "karl.moberg@evry.com",
		"source_code" : "https://github.com/kmoberg/gopher",
	})

	// Do error handling, just in case.
	if err != nil {
		log.Fatal(err)
	}

	// Post
	resp, err := http.Post("http://10.200.233.89:8080/contest", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	// Get response
	body, err := ioutil.ReadAll(resp.Body)

	// More error handling. 
	if err != nil {
		log.Fatalln(err)
	}

	// Print the response
	log.Println(string(body))


}
