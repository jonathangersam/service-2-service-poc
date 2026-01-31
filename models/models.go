package models

type Comment struct {
	CommentId      string
	GlobalEntityId string
	OrderId        string
	Text           string
}

type Order struct {
	OrderId        string
	GlobalEntityId string // pretend this is something only OrderAPI knows
	Status         string
}
