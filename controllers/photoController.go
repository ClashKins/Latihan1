package controllers

import (
	"LATIHAN1/database"
	"LATIHAN1/helpers"
	"LATIHAN1/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))
	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	}else {
		c.ShouldBind(&Photo)
	}
	Photo.UserID = userID

	Result := map[string]interface{}{}

	 err := db.Raw(
		"INSERT into photos (title, caption, photo_url, user_id, created_at) VALUES(?,?,?,?,?) Returning id,title,photo_url, user_id, created_at, caption",
		Photo.Title, Photo.Caption, Photo.PhotoUrl, Photo.UserID, time.Now(),
	 ).Scan(&Result).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Result)
	}

func UpdatePhoto(c *gin.Context){
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		contentType := helpers.GetContentType(c)
		_, _ = db, contentType
		Photo := models.Photo{}
		photoId, _ := strconv.Atoi(c.Param("photoId"))
		userID := uint(userData["id"].(float64))
		if contentType == appJSON{
			c.ShouldBindJSON(&Photo)
		}else{
			c.ShouldBind(&Photo)
		}
		Photo.UserID = userID

		Result := map[string]interface{}{}
		SqlStatement := "Update photos SET title = ?, caption = ?, photo_url = ?, updated_at = ? WHERE id = ? RETURNING id, title, caption, photo_url, updated_at"

		err := db.Raw(
			SqlStatement,
			Photo.Title, Photo.Caption, Photo.PhotoUrl, time.Now(), uint(photoId),
		).Scan(&Result).Error
	
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "Bad Request",
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, Result)
	}

	func GetAllPhoto(c *gin.Context) {
		db := database.GetDB()
		Photos := []models.Photo{}
		err :=  db.Debug().Preload("User").Preload("Comments").Find(&Photos).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"err": "Bad Request",
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, Photos)
	}

	func GetOnePhoto(c *gin.Context) {
		db := database.GetDB()
		photoId, _ := strconv.Atoi(c.Param("photoId"))
		Photos := models.Photo{}
		err := db.Preload("User").Find(&Photos, uint(photoId)).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "Bad Request",
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, Photos)
	}

	func DeletePhoto(c *gin.Context) {
		db := database.GetDB()
		photoId, _ := strconv.Atoi(c.Param("photoId"))
		Photo := models.Photo{}
		err := db.Delete(Photo, uint(photoId)).Error

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "Bad Request",
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Your photo has been successfully deleted",
		})
	}
