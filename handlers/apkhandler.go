package handlers

import (
	"net/http"

	"github.com/ALZEE23/ApiGo/database"
	"github.com/ALZEE23/ApiGo/models"
	"github.com/gin-gonic/gin"
)

func Apk(context *gin.Context) {
	var apk models.Apk
	if err := context.ShouldBindJSON(&apk); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
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
