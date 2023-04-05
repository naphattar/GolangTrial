package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/naphattar/KaihorBackend/configs"
	"github.com/naphattar/KaihorBackend/models"
	"github.com/naphattar/KaihorBackend/responses"
	"github.com/naphattar/KaihorBackend/utills"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var campCollection *mongo.Collection = configs.GetCollection(configs.DB, "camps")

func UpdateCampDataFromSpreadSheet(c *fiber.Ctx) error {
	campsData := utills.GetDatafromSheet()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// delete all old Data in the Database
	filter := bson.D{{}}
	_, err := campCollection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(responses.UserResponse{Status: http.StatusBadGateway, Message: "error", Data: &fiber.Map{"error": err.Error()}})
	}

	// add new Data the the Database
	for i := 1; i < len(campsData); i++ {
		campData := campsData[i]
		newCampData := models.Camp{
			CampID:   campData.CampID,
			Name:     campData.Name,
			Time:     campData.Time,
			Location: campData.Location,
			Director: campData.Director,
		}
		_, err := campCollection.InsertOne(ctx, newCampData)
		if err != nil {
			return c.Status(http.StatusBadGateway).JSON(responses.UserResponse{Status: http.StatusBadGateway, Message: "error", Data: &fiber.Map{"error": err.Error()}})
		}
	}
	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "data updated success"})
}
