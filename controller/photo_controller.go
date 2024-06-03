package controller

import (
	"encoding/json"
	"final-task-pbi-rakamin-fullstack-muhammad-rifqi-setiawan/database"
	"final-task-pbi-rakamin-fullstack-muhammad-rifqi-setiawan/helper"
	"final-task-pbi-rakamin-fullstack-muhammad-rifqi-setiawan/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userInfo").(jwt.MapClaims)
	contentType := helper.GetContentType(ctx)

	photoRequest := model.CreatePhotoRequest{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		if err := ctx.ShouldBindJSON(&photoRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, helper.CreateResponse(false, nil, err.Error(), "Invalid request"))
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, helper.CreateResponse(false, nil, "Invalid content type", "Content type must be application/json"))
		return
	}

	photo := model.Photo{
		Title:    photoRequest.Title,
		Caption:  photoRequest.Caption,
		PhotoURL: photoRequest.PhotoURL,
		UserID:   userID,
	}

	err := db.Debug().Create(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error(), "Failed to create photo"))
		return
	}
	_ = db.First(&photo, photo.ID).Error

	photoString, _ := json.Marshal(photo)
	photoResponse := model.CreatePhotoResponse{}
	json.Unmarshal(photoString, &photoResponse)

	ctx.JSON(http.StatusCreated, helper.CreateResponse(true, photoResponse, "", "Photo created successfully"))
}

func GetPhotoById(ctx *gin.Context) {
	db := database.GetDB()
	PhotoID := ctx.Param("PhotoId")
	photo := model.Photo{}

	err := db.Debug().Preload("User").First(&photo, PhotoID).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.CreateResponse(false, nil, err.Error(), "Photo not found"))
		return
	}

	photoString, _ := json.Marshal(photo)
	photoResponse := model.GetPhotoResponse{}
	json.Unmarshal(photoString, &photoResponse)

	ctx.JSON(http.StatusOK, helper.CreateResponse(true, photoResponse, "", "Photo fetched successfully"))
}

func UpdatePhoto(ctx *gin.Context){
	db := database.GetDB()
	userData := ctx.MustGet("userInfo").(jwt.MapClaims)
	contentType := helper.GetContentType(ctx)

	photoRequest := model.UpdatePhotoRequest{}
	photoId, _ := strconv.Atoi(ctx.Param("PhotoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		// Bind JSON data to UpdatePhotoRequest
		if err := ctx.ShouldBindJSON(&photoRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, helper.CreateResponse(false, nil, err.Error(), "Invalid request"))
			return
		}
	} else {
		 // If content type is not JSON, respond with error
		 ctx.JSON(http.StatusBadRequest, helper.CreateResponse(false, nil, "Invalid content type", "Content type must be application/json"))
		 return
	}

	photo := model.Photo{}
	photo.ID = photoId
	photo.UserID = userID

	// Check if photo exists
	if err := db.First(&photo, photoId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, helper.CreateResponse(false, nil, err.Error(), "Photo not found"))
		return
	}

	// check if user is authorized to update photo
	if photo.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "Unauthorized", "You are not authorized to update this photo"))
		return
	}

	// marshal photoRequest to string
	updateString, _ := json.Marshal(photoRequest)
	updateData := model.Photo{}
	json.Unmarshal(updateString, &updateData)

	// Perform update
	if err := db.Model(&photo).Updates(updateData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error(), "Failed to update photo"))
		return
	}

	// Fetch updated photo
	updatedPhoto := model.Photo{}
	if err := db.First(&updatedPhoto, photoId).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error(), "Failed to fetch updated photo"))
		return
	}

	photoString , _ := json.Marshal(updatedPhoto)
	photoResponse := model.UpdatePhotoResponse{}
	json.Unmarshal(photoString, &photoResponse)

	ctx.JSON(http.StatusOK, helper.CreateResponse(true, photoResponse, "", "Photo updated successfully"))
}

func DeletePhoto(ctx *gin.Context) {
	database := database.GetDB()
	userData := ctx.MustGet("userInfo").(jwt.MapClaims)

	photoID, err := strconv.Atoi(ctx.Param("PhotoId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.CreateResponse(false, nil, err.Error(), "Invalid photo ID"))
		return
	}

	userID := uint(userData["id"].(float64))

	photo := model.Photo{}
	if err := database.First(&photo, photoID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, helper.CreateResponse(false, nil, err.Error(), "Photo not found"))
		return
	}

	if photo.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "Unauthorized", "You are not authorized to delete this photo"))
		return
	}

	if err := database.Delete(&photo).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error(), "Failed to delete photo"))
		return
	}
	
    ctx.JSON(http.StatusOK, helper.CreateResponse(true, nil, "", "Your photo has been successfully deleted "))
}