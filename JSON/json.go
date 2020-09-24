package main

import (
	"encoding/json"
	"fmt"

	"os"
)

type Post struct {
	Id 	int `json: "id"`
	Content string `json: "content"`
	Author Author `json: "author"`
	Comments []Comment `json: "comments"`
}

type Author struct {
	Id int `json: "id"`
	Name string `json: "name"`
}

type Comment struct {
	Id int `json: "id"`
	Content string `json: "content"`
	Author string `json: "author"`
}

func main() {
	post := Post{
		Id : 1,
		Content: "HEllo world",
		Author: Author{
			Id: 2,
			Name: "Sau Sheong",
		},
		Comments: []Comment{
			Comment{
			Id: 3,
			Content: "Hav a great day",
			Author: "Admam",
		},
		Comment{
			Id: 4,
			Content: "How ary you today",
		},
	},
}

jsonFile, err := os.Create("post.json")
if err != nil {
	fmt.Println("Error creating json:", err)
	return
}
encoder := json.NewEncoder(jsonFile)
err = encoder.Encode(&post)
if err != nil {
	fmt.Println("Error encoding json to file:", err)
	return
}
}