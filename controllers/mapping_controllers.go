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
	"go.mongodb.org/mongo-driver/mongo"
)

func mappingCollection() *mongo.Collection {
	return config.DB.Collection("mappings")
}

func CreateMapping(c *gin.Context) {
	var mappings []models.Mapping
	if err := c.ShouldBindJSON(&mappings); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid request",
			"detail": err.Error(),
		})
		return
	}

	for i := range mappings {
		mappings[i].ID = primitive.NewObjectID()
		mappings[i].CreatedAt = time.Now()
		mappings[i].UpdatedAt = time.Now()

		_, err := mappingCollection().InsertOne(context.Background(), mappings[i])
		if err != nil {
			helper.ErrorResponse(c, http.StatusInternalServerError, "Terjadi kesalahan", err.Error())
			return
		}
	}

	helper.ResponseSucces(c, http.StatusOK, "Berhasil membuat mapping", mappings)
}

func GetMappings(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := mappingCollection().Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	var mappings []models.Mapping = []models.Mapping{}
	if err = cursor.All(ctx, &mappings); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mappings)
}

func UpdateMapping(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "ID tidak valid", err.Error())
		return
	}

	var updateData models.Mapping
	if err := c.ShouldBindJSON(&updateData); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Terjadi kesalahan", err.Error())
		return
	}
	updateData.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"kesatuan":  updateData.Kesatuan,
			"angkatan":  updateData.Angkatan,
			"updatedAt": updateData.UpdatedAt,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := mappingCollection().UpdateByID(ctx, objID, update)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Gagal update mapping", err.Error())
		return
	}

	helper.ResponseSucces(c, http.StatusOK, "Berhasil update mapping", result)
}

func DeleteMapping(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "ID tidak valid", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := mappingCollection().DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Gagal hapus mapping", err.Error())
		return
	}

	if result.DeletedCount == 0 {
		helper.ErrorResponse(c, http.StatusNotFound, "Mapping tidak ditemukan", "")
		return
	}

	helper.ResponseSucces(c, http.StatusOK, "Berhasil hapus mapping", result)
}

func DeleteSemuaMapping(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := mappingCollection().DeleteMany(ctx, bson.M{})
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Gagal hapus semua mapping", err.Error())
		return
	}

	if result.DeletedCount == 0 {
		helper.ErrorResponse(c, http.StatusNotFound, "Tidak ada mapping yang ditemukan untuk dihapus", "")
		return
	}

	helper.ResponseSucces(c, http.StatusOK, "Berhasil hapus semua mapping", gin.H{
		"deletedCount": result.DeletedCount,
	})
}
