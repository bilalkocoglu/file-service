package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imminoglobulin/file-service/pkg/minio"
	"net/http"
)

func UploadFile(ctx *gin.Context) {
	filename := ctx.PostForm("filename")
	file, err := ctx.FormFile("file")
	cfg := minio.BuildMinioConfig()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("get form err: %s", err.Error()),
		})
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "File error",
		})
	}
	uploadInfo, err := minio.UploadFile(cfg.MainBucketName, filename, openedFile, file.Size)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("get form err: %s", err.Error()),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "SUCCESS",
			"info":    uploadInfo,
		})
	}
}
