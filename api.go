package main

type bookmarkDto struct {
	Page int `json:"page" binding:"required,gte=1"`
}
