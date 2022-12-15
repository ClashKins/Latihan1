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

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType :=  helpers.GetContentType(c)
	_, _ = db, contentType
	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))
	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	}else {
		c.ShouldBind(&SocialMedia)
	}
	SocialMedia.UserID = userID

	Result := map[string]interface{}{}

	err := db.Raw(
		`INSERT into social_media
		(name, social_media_url, user_id, created_at) VALUES(?,?,?,?)
		Returning id, name, social_media_url, user_id, created_at`,
		SocialMedia.Name, SocialMedia.SocialMediaUrl, SocialMedia.UserID, time.Now(),
	).Scan(&Result).Error

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"err": "Bad Request",
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, Result)
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_,_ = db, contentType
	SocialMedia := models.SocialMedia{}
	socialmediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))
	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	}else {
		c.ShouldBind(&SocialMedia)
	}
	SocialMedia.UserID = userID

	Result := map[string]interface{}{}
	err := db.Raw(
		`Update social_media
		SET name = ?, social_media_url = ?, updated_at = ?
		WHERE = ?
		Returning id, name, social_media_url, updated_at`,
		SocialMedia.Name, SocialMedia.SocialMediaUrl, time.Now(), uint(socialmediaId),
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

func GetAllSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedias := []models.SocialMedia{}
	err := db.Debug().Preload("User").Preload("SocialMedias").Find(&SocialMedias).Error

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, SocialMedias)
}

func GetOneSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}
	socialMediaId,_ := strconv.Atoi(c.Param("socialMediaId"))
	err := db.Debug().Preload("User").Preload("SocialMedia").Find(&SocialMedia, socialMediaId).Error
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, SocialMedia)
}

func DeleteSocialMedia(c *gin.Context){
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	err := db.Delete(SocialMedia, uint(socialMediaId)).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
