package models

import "time"

type Task struct {
	Id        int       `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	IsDeleted bool      `json:"-"`
	Create_at time.Time `json:"-"`
	Update_at time.Time `json:"-"`
}
