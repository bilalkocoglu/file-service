package api

import (
	"github.com/gin-gonic/gin"
	"github.com/imminoglobulin/file-service/pkg/database"
	"net/http"
	"strconv"
)

func AllUsers(ctx *gin.Context) {
	var AppUsers []database.ApplicationUser

	database.GetAllAppUsers(&AppUsers)

	ctx.JSON(http.StatusOK, AppUsers)
}

func SaveUser(ctx *gin.Context) {
	var appUser database.ApplicationUser
	if err := ctx.ShouldBindJSON(&appUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Invalid json provided",
		})
		return
	}

	err := database.SaveUser(&appUser)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
}

func FindUserById(ctx *gin.Context) {
	var appUser database.ApplicationUser
	id := ctx.Param("id")
	intId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Id must be numeric value",
		})
		return
	}
	database.GetUserById(&appUser, intId)

	if appUser.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
	} else {
		ctx.JSON(http.StatusOK, appUser)
	}
}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Id must be numeric value",
		})
		return
	}
	database.DeleteUserById(intId)
}
