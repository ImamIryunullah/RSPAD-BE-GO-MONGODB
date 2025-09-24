package controllers

import (
	"context"
	"net/http"
	"time"

	"spad_be/config"
	helper "spad_be/helpers"
	"spad_be/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create User
func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		helper.ErrorResponse(c, http.StatusBadRequest, "error", "")
		return
	}

	collection := config.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Gagal insert user!", err.Error())
		return

	}
	helper.ResponseSucces(c, http.StatusOK, "User berhasil dibuat", res)
}

// Get All Users
func GetUsers(c *gin.Context) {
	collection := config.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Gagal ambil user", "")
		return
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Gagal decode data!", err.Error())
		return
	}

	helper.ResponseSucces(c, http.StatusOK, "berhasil mengambil data Users", users)
}

// Get User by ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "ID tidak valid", "ID tidak valid")
		return
	}

	collection := config.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "User tidak ditemukan", "")
		return
	}

	helper.ResponseSucces(c, http.StatusOK, "berhasil mengambil data Users by ID", user)
}
