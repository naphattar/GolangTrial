package controllers

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/naphattar/KaihorBackend/models"
	"github.com/naphattar/KaihorBackend/responses"
	scrappers "github.com/naphattar/KaihorBackend/utills"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllCampData(c *fiber.Ctx) error {
	filter := bson.D{{}}
	campsData, err := campCollection.Find(context.TODO(), filter)
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(responses.UserResponse{Status: http.StatusBadGateway, Message: "error", Data: &fiber.Map{"error": err.Error()}})
	}
	var results []bson.M
	// check for errors in the conversion
	if err = campsData.All(context.TODO(), &results); err != nil {
		return c.Status(http.StatusBadGateway).JSON(responses.UserResponse{Status: http.StatusBadGateway, Message: "error", Data: &fiber.Map{"error": err.Error()}})
	}
	return c.Status(http.StatusAccepted).JSON(responses.UserResponse{Status: http.StatusAccepted, Message: "success", Data: &fiber.Map{"camps": results}})
}

func GetCampDatabyID(c *fiber.Ctx) error {
	id := c.Params("id")
	var campData models.Camp
	filter := bson.D{{"campid", id}}
	err := campCollection.FindOne(context.TODO(), filter).Decode(&campData)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{Status: http.StatusNotFound, Message: "campid not found", Data: &fiber.Map{"error": err.Error()}})
	}
	//var results []bson.M
	return c.Status(http.StatusAccepted).JSON(responses.UserResponse{
		Status:  http.StatusAccepted,
		Message: "success",
		Data:    &fiber.Map{"camps": campData}})
}

func GetCampDatabyLocation(c *fiber.Ctx) error {
	location := c.Params("location")
	location, err := scrappers.DecodeUrl(location)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "location is invalid",
			Data:    &fiber.Map{"data": err.Error()}})
	}

	filter := bson.D{{"location", location}}
	campsData, err := campCollection.Find(context.TODO(), filter)

	var results []bson.M

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{
			Status:  http.StatusNotFound,
			Message: "camp in this location not found",
			Data:    &fiber.Map{"data": err.Error()}})
	}
	if err = campsData.All(context.TODO(), &results); err != nil {
		return c.Status(http.StatusBadGateway).JSON(responses.UserResponse{Status: http.StatusBadGateway, Message: "error", Data: &fiber.Map{"error": err.Error()}})
	}
	return c.Status(http.StatusAccepted).JSON(responses.UserResponse{Status: http.StatusAccepted, Message: "success", Data: &fiber.Map{"camps": results}})
}
