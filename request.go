package main

import "net/http"
import "encoding/json"

func doRequest(endpoint string, object interface{}) {

	url := "https://baseurl/" + endpoint
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "token")
	request.Header.Set("Content-Type", "application/json")	

	response, _ := client.Do(request)
	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(object)
}