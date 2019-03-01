package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Post struct {
	Id         int
	Content    string
	AuthorName string `db: author`
}

var Db *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	checkError(err)
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRowx("select id, content, author from posts where id = $1", id).StructScan(&post)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (post *Post) Create() (err error) {
	err = Db.QueryRow("insert into posts (content, author) values ($1,$2) returning id", post.Content, post.AuthorName).Scan(&post.Id)
	checkError(err)
	return
}

func checkError(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post{Content: "Hello World!", AuthorName: "kido_k"}
	post.Create()
	fmt.Println(post)

	readPost := Post{}
	readPost, _ = GetPost(post.Id)
	fmt.Println(readPost)
}
