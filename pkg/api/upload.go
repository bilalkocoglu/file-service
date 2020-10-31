package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadFile(ctx *gin.Context) {
	filename := ctx.PostForm("filename")
	file, err := ctx.FormFile("file")

	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	if err := ctx.SaveUploadedFile(file, "upload/"+filename); err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success !",
	})
}
