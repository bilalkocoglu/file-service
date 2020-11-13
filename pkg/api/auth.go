package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_const "github.com/imminoglobulin/file-service/pkg/const"
	"github.com/imminoglobulin/file-service/pkg/database"
	"github.com/imminoglobulin/file-service/pkg/model"
	"net/http"
	"os"
	"time"
)

func Login(c *gin.Context) {
	var dto model.Login
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Invalid json provided",
		})
		return
	}

	var user database.ApplicationUser
	err := database.GetUserByUsername(&user, dto.Username)

	if err != nil || user.ID == 0 || user.Password != dto.Password {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username and password.",
		})
		return
	}

	token, err := CreateToken(user.ID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Can not token create",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": "Bearer " + token,
	})
}

func CreateToken(id uint) (string, error) {
	var err error
	os.Setenv("ACCESS_SECRET", _const.SecretKey)
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(_const.TokenExpireTime).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
