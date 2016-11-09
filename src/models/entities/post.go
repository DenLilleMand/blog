package entities

import (
	"time"
)

type Visibility int

const (
	Public  Visibility = iota + 1
	Private            = iota
)

var visibilities = [...]string{
	"Public",
	"Private",
}

func (Visibility Visibility) String() string {
	return visibilities[Visibility]
}

type Post struct {
	Id         int
	Author     int
	Created    time.Time
	Title      string
	Visibility int
	Updated    time.Time
	Guid       string
}
