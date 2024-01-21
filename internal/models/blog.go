package models

import (
	"time"
)

type Blog struct {
	ID        int64     `json:"id" db:"id" example:"1"`
	Title     string    `json:"title" db:"title" validate:"required,gte=3" example:"this is title"`
	Content   string    `json:"content" db:"content" validate:"required,gte=10" example:"this is content"`
	CreatedAt time.Time `json:"created_at" db:"created_at" example:"2021-01-01T00:00:00Z"`
}

type BlogList struct {
	TotalCount int     `json:"total_count" example:"100"`
	TotalPage  int     `json:"total_page" example:"10"`
	Page       int     `json:"page" example:"1"`
	Limit      int     `json:"limit" example:"10"`
	HasMore    bool    `json:"has_more" example:"true"`
	Blogs      []*Blog `json:"blogs" example:[{id:1,title:this is title,content:this is content,created_at:2021-01-01T00:00:00Z}]`
}

type BlogSwagger struct {
	Title   string `json:"title" validate:"required,gte=3,max=255" example:"this is title"`
	Content string `json:"content" validate:"required,gte=10" example:"this is content"`
}
