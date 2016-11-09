package repositories

import (
	"database/sql"
	"fmt"
	"github.com/denlillemand/blog/src/models/entities"
)
import _ "github.com/go-sql-driver/mysql"

type PostRepository struct{}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func (postRepository *PostRepository) Get(id int) (post *entities.Post, err error) {
	if id < 1 {
		return nil, err
	}
	db, err := sql.Open("mysql", "root:denlilleiceman20@/codeandcocaine?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_ = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("repository get: %v", id)
	fetchedPost := new(entities.Post)
	err = db.QueryRow("SELECT post_id, author, created, title, guid, updated FROM posts WHERE post_id=?", id).Scan(&fetchedPost.Id, &fetchedPost.Author, &fetchedPost.Created, &fetchedPost.Title, &fetchedPost.Guid, &fetchedPost.Updated)
	fmt.Printf("the fetched post: %v", fetchedPost)
	switch {
	case err == sql.ErrNoRows:
		fmt.Printf("No post with that id")
		return nil, err
	case err != nil:
		fmt.Errorf("A error happended: %v", err)
		return nil, err
	default:
		return fetchedPost, nil
	}
}
