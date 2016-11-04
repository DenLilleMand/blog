package repositories

import (
	"database/sql"
	"github.com/denlillemand/blog/models/entities"
	"time"
)

type PostRepository struct {
}

func (postRepository PostRepository) get(id int) Post {
	return Post{
		Id:           1,
		PostAuthor:   1,
		PostDate:     time.Now(),
		PostTitle:    "herpderp",
		Visibility:   0,
		PostModified: time.Now(),
		Guid:         "223232-42424-2424-2424",
	}
}
