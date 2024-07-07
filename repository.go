package main

import (
	"errors"
	"log/slog"

	"gorm.io/gorm"
)

func readCurrentPage(config *config) int {
	var storedBookmark bookmark
	if err := config.Database.First(&storedBookmark, 1).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return 1
		} else {
			slog.Error("Something went wrong", "error", err.Error())
			return 1
		}
	}
	return storedBookmark.Page
}

func updateCurrentPage(config *config, newPage int) int {
	var storedBookmark bookmark
	if err := config.Database.First(&storedBookmark).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			config.Database.Create(&bookmark{Page: newPage})
			return newPage
		} else {
			slog.Error("something went wrong", "error", err.Error())
			return 1
		}
	}

	config.Database.Model(&storedBookmark).Update("Page", newPage)
	return newPage
}
