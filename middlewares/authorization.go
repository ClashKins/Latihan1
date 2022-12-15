package middlewares

import (
	"LATIHAN1/database"
	"LATIHAN1/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		photoId, err := strconv.Atoi(c.Param("photoId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := models.User{}
		userLevel := db.Select("level").First(&User, "id = ?", uint(userID))
		_ = userLevel

		Photo := models.Photo{}
		err = db.Select("user_id").First(&Photo, uint(photoId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Data Not Found!",
				"message": "data doesn't exist",
			})
			return
		}
		if User.Level == "admin" || User.Level == "superadmin" {
			return
		} else if User.Level == "user" {
			if Photo.UserID != userID {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Unauthorized",
					"message": "you are not allowed to access this data",
				})
				return
			}
		}else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		commentId, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad request",
				"message": "invalid parameter",
			})
			return
		}
		userData := c.MustGet("userdata").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Comment := models.Comment{}
		err = db.Select("user_id").First(&Comment, uint(commentId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}
		if Comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"message": "data doesn't exist",
			})
			return
		}
		if Comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		socialMediaId, err := strconv.Atoi(c.Param("socialmediaId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad request",
				"message": "invalid parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		SocialMedia := models.SocialMedia{}

		err = db.Select("user_id").First(&SocialMedia, uint(socialMediaId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Data not found",
				"message": "data doesn't exist",
			})
			return
		}
		if SocialMedia.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}