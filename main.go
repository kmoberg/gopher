package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"log"
)


func main() {

	requestBody, err := json.Marshal(map[string]string{
		"name" : "Karl Mathias Moberg",
		"email" : "karl.moberg@evry.com",
		"source_code" : "https://github.com/kmoberg/gopher",
	})

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://10.200.233.89:8080/contest", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))


}
