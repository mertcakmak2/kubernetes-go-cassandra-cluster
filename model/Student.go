package model

import "time"

type Student struct {
	ID        int       `json:"ID"`
	Firstname string    `json:"Firstname"`
	Lastname  string    `json:"Lastname"`
	Age       int       `json:"Age"`
	Lastdate  time.Time `json:"Lastdate"`
}
