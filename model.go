package main

import "gorm.io/gorm"

type bookmark struct {
	gorm.Model
	Page int
}
