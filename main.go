package main

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("bookmark.db"))
	if err != nil {
		return
	}

	db.AutoMigrate(&bookmark{})

	config := &config{Database: db}

	r := gin.New()
	r.GET("/currentPage", getBookmark(config))
	r.POST("/currentPage", setBookmark(config))
	r.Run()
}

func getBookmark(config *config) func(*gin.Context) {
	return func(c *gin.Context) {
		currentPage := readCurrentPage(config)
		c.JSONP(http.StatusOK, bookmarkDto{currentPage})
	}
}

func setBookmark(config *config) func(c *gin.Context) {
	return func(c *gin.Context) {
		var newBookmark bookmarkDto

		if err := c.BindJSON(&newBookmark); err != nil {
			slog.Error(err.Error())
			c.Status(http.StatusBadRequest)
			return
		}

		storedPage := updateCurrentPage(config, newBookmark.Page)
		c.JSONP(http.StatusOK, bookmarkDto{storedPage})
	}
}
