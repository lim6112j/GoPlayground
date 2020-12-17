package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func PrintJsonResponse(resp *http.Response) {
	if resp.Body == nil {
		return
	}
	decoder := json.NewDecoder(resp.Body)
	var posts map[string]interface{}
	err := decoder.Decode(&posts)
	if err != nil {
		log.Printf("Error while print json decoding : %v", err)
	}
	b, err := json.Marshal(posts)
	if err != nil {
		log.Printf("Error while print json marshalling : %v", err)
		return
	}
	fmt.Println(string(b))
}
