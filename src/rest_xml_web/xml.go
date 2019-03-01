package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Post struct {
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  string `xml:"author"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	//xml decodeで要素毎の解析
	decoder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se)
				fmt.Println("comment\n", comment)
			}
		}
	}

	// //xml全文解析
	// xmlData, err := ioutil.ReadAll(xmlFile)
	// if err != nil {
	// 	fmt.Println("Error reading XML data:", err)
	// 	return
	// }
	// var post Post
	// xml.Unmarshal(xmlData, &post)

	// // fmt.Println(post)
	// fmt.Println("Id\n", post.Id)
	// fmt.Println("Content\n", post.Content)
	// fmt.Println("Author\n", post.Author)
	// fmt.Println("Comments\n", post.Comments)
}
