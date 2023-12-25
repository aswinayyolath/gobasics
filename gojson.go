package main

import (
	"encoding/json"
	"fmt"
)

// Course represents a course with JSON tags for encoding
type Course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Password string   `json:"-"` // Password field will be removed from the JSON
	Platform string   `json:"platform"`
	Tag      []string `json:"tag,omitempty"` // Tag will be removed from JSON if it's empty
}

func main() {
	fmt.Println("Working with JSON in Golang")
	// encodeJSON()
	DecodeJson()
}

// encodeJSON encodes a slice of Course structs into JSON format with indentation
func encodeJSON() {
	// Sample data of courses
	data := []Course{
		{
			"ReactJs",
			1000,
			"abc123",
			"Online",
			[]string{"web-development", "js"},
		},
		{
			"NodeJs",
			400,
			"bcd123",
			"Offline",
			[]string{"Backend-development", "js"},
		},
		{
			"Python",
			1000,
			"efg123",
			"Online",
			nil,
		},
	}

	// MarshalIndent converts the data into a formatted JSON string
	finalJSON, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		// Handle the error if there's an issue with JSON encoding
		panic(err)
	}

	// Print the formatted JSON
	fmt.Printf("%s", finalJSON)
}

func DecodeJson() {
	jsonData := []byte(`

	{
		"coursename": "ReactJs",
		"price": 1000,
		"platform": "Online",
		"tag": [
				"web-development",
				"js"
		]
	}
	`)
	var courseData Course
	if json.Valid(jsonData) {
		json.Unmarshal(jsonData, &courseData)
		fmt.Printf("%#v\n", courseData)
	} else {
		fmt.Println("JSON is invalid")
	}

	// In some cases we just want to add data to key and value

	var keyValueData map[string]interface{}
	json.Unmarshal(jsonData, &keyValueData)
	for k, v := range keyValueData {
		fmt.Printf("Key: %v  Value: %v and Type %T\n", k, v, v)
	}
}
