package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/ALZEE23/ApiGo/database"
	"github.com/ALZEE23/ApiGo/models"
	"github.com/gin-gonic/gin"
)

func Apk(context *gin.Context) {
	_, err := context.MultipartForm()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form data"})
		return
	}

	footageHeader, err := context.FormFile("footage")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get footage file"})
		return
	}

	coverHeader, err := context.FormFile("cover")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get cover file"})
		return
	}

	footagePath := filepath.Join("storage", footageHeader.Filename)
	if err := context.SaveUploadedFile(footageHeader, footagePath); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save footage file"})
		return
	}

	coverPath := filepath.Join("storage", coverHeader.Filename)
	if err := context.SaveUploadedFile(coverHeader, coverPath); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save cover file"})
		return
	}

	var apk models.Apk
	if err := context.ShouldBindJSON(&apk); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	apk.Footage = footagePath
	apk.Cover = coverPath

	record := database.DB.Db.Create(&apk)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"apkId": apk.ID, "name": apk.Name, "game": apk.Game, "cover": apk.Cover, "title": apk.Title, "description": apk.Description, "footage": apk.Footage})
}

var retData struct {
	List []models.Apk `json:"list"`
}

func GetApk(context *gin.Context) {
	apks := database.DB.Db.Find(&retData.List)
	if apks.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": apks.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": retData.List})
}
