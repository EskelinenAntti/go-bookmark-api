package main

import "gorm.io/gorm"

type config struct {
	Database *gorm.DB
}
