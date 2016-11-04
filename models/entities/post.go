package entities

import (
	"fmt"
	"time"
)

type Visibility int

const (
	Public  Visibility = iota + 1 //1
	Private            = iota     //2
)

var visibilities = [...]string{
	"Public",
	"Private",
}

func (Visibility Visibility) String() string {
	return visibilities[Visibility]
}

type Post struct {
	Id           int
	PostAuthor   int
	PostDate     time
	PostTitle    string
	Visibility   int
	PostModified time
	Guid         string
}
