package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Body  string
}

//gorm.Model is special field which includes common columns like Id, CreatedAt, UpdatedAt and DeletedAt
