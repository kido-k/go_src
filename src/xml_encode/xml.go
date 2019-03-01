package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post{
		Id:      "1",
		Content: "Hello World!",
		Author: Author{
			Id:   "2",
			Name: "kido_k",
		},
	}

	// output, err := xml.Marshal(&post)
	// if err != nil {
	// 	fmt.Println("Error opening XML file:", err)
	// 	return
	// }
	// err = ioutil.WriteFile("post1.xml", output, 0644)
	// if err != nil {
	// 	fmt.Println("Error writing XML to file:", err)
	// 	return
	// }

	xmlFile, err := os.Create("post2.xml")
	if err != nil {
		fmt.Println("Error creating XML file:", err)
		return
	}
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding XML to file:", err)
	}
}
