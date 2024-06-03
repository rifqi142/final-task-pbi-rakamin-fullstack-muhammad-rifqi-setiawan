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

var (
	appJSON = "application/json"
)

// Register User
func Register(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(ctx)

	userRequest := model.CreateUserRequest{}

	// Check Content Type
	if contentType == appJSON {
        // Bind JSON data to CreateUserRequest
        if err := ctx.ShouldBindJSON(&userRequest); err != nil {
            ctx.JSON(http.StatusBadRequest, helper.CreateResponse(false, nil, err.Error(), "Invalid request"))
            return
        }
    } else {
        // If content type is not JSON, respond with error
        ctx.JSON(http.StatusBadRequest, helper.CreateResponse(false, nil, "Invalid content type", "Content type must be application/json"))
        return
    }

	user := model.User {
		Username: userRequest.Username,
		Email: userRequest.Email,
		Password: helper.HashPassword(userRequest.Password),
	}

	err := db.Debug().Create(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error(), "Failed to create user"))
		return
	}

	userString, _ := json.Marshal(user)
	userResponse := model.CreateUserResponse{}
	json.Unmarshal(userString, &userResponse)

	ctx.JSON(http.StatusCreated, helper.CreateResponse(true, userResponse, "", "User created successfully"))
}

// Login user
func Login(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(ctx)

	loginRequest := model.LoginUserRequest{}

	if contentType == appJSON {
		if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, helper.CreateResponse(false, nil, err.Error(), "Invalid request"))
			return
		}
	} else {
		if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, helper.CreateResponse(false, nil, err.Error(), "Invalid request"))
			return
		}
	}

	password := loginRequest.Password
	user := model.User{}

	err := db.Debug().Where("email = ?", loginRequest.Email).First(&user).Error
	// check if user not found
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.CreateResponse(false, nil, err.Error(), "User not found"))
		return
	}

	// check if password is invalid
	if !helper.CheckPasswordHash(password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "", "Invalid password"))
		return
	}

	// generate token
	token := helper.GenerateToken(user.ID, user.Email)

	// return token
	ctx.JSON(http.StatusOK, helper.CreateResponse(true, map[string]interface{}{"Token": token}, "", "Login success"))
}

func UpdateDataUser(ctx *gin.Context) {
    db := database.GetDB()
    userData := ctx.MustGet("userInfo").(jwt.MapClaims)
    contentType := helper.GetContentType(ctx)

    userRequest := model.UpdateUserRequest{}
    userID := uint(userData["id"].(float64))

    // Check if content type is JSON
    if contentType != appJSON {
        ctx.JSON(http.StatusBadRequest, helper.CreateResponse(false, nil, "Invalid content type", "Content type must be application/json"))
        return
    }

    // Bind JSON data to CreateUserRequest
    if err := ctx.ShouldBindJSON(&userRequest); err != nil {
        ctx.JSON(http.StatusBadRequest, helper.CreateResponse(false, nil, err.Error(), "Invalid request"))
        return
    }

    user := model.User{}
    user.ID = userID

    // Check if user exists
    if err := db.First(&user, userID).Error; err != nil {
        ctx.JSON(http.StatusNotFound, helper.CreateResponse(false, nil, "User not found", "User does not exist in the database"))
        return
    }

    // Marshal userRequest to string
    updateString, _ := json.Marshal(userRequest)
    updateData := model.User{}
    json.Unmarshal(updateString, &updateData)

    // Perform update
    if err := db.Model(&user).Updates(updateData).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error(), "Failed to update user"))
        return
    }

    // Fetch updated user
    updatedUser := model.User{}
    if err := db.First(&updatedUser, userID).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error(), "Failed to fetch updated user"))
        return
    }

    userString, _ := json.Marshal(updatedUser)
    userResponse := model.UpdateUserResponse{}
    json.Unmarshal(userString, &userResponse)

    ctx.JSON(http.StatusOK, helper.CreateResponse(true, userResponse, "", "User updated successfully"))
}


func DeleteUser(ctx *gin.Context) {
    db := database.GetDB()
    
    userIDStr := ctx.Param("UserId")
    userID, err := strconv.ParseUint(userIDStr, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, helper.CreateResponse(false, nil, "Invalid user ID", "User ID must be a number"))
        return
    }

    user := model.User{}

    if err := db.First(&user, userID).Error; err != nil {
        ctx.JSON(http.StatusNotFound, helper.CreateResponse(false, nil, "User not found", "User does not exist in the database"))
        return
    }

    if err := db.Delete(&user).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error(), "Failed to delete user"))
        return
    }

    ctx.JSON(http.StatusOK, helper.CreateResponse(true, nil, "", "Your account has been successfully deleted"))
}
