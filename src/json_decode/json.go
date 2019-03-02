package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id,attr"`
	Name string `json:",chardata"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	//read all json file
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}

	var post Post
	json.Unmarshal(jsonData, &post)
	// fmt.Println(post)
	fmt.Println("Id: ", post.Id)
	fmt.Println("Content: ", post.Content)
	fmt.Println("Author: ", post.Author)
	fmt.Println("Comments: ", post.Comments)

	// read all json file
	// decoder := json.NewDecoder(jsonFile)
	// for {
	// 	var post Post
	// 	err := decoder.Decode(&post)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("Error decoding JSON:", err)
	// 		return
	// 	}
	// 	fmt.Println(post)
	// }
}
