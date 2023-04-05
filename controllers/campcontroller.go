package controllers

import (
	"context"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/naphattar/KaihorBackend/models"
	"github.com/naphattar/KaihorBackend/responses"
	"github.com/naphattar/KaihorBackend/utills"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllCampData(c *fiber.Ctx) error {
	filter := bson.D{{}}
	campsData, err := campCollection.Find(context.TODO(), filter)
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(responses.UserResponse{Status: http.StatusBadGateway, Message: "error", Data: &fiber.Map{"error": err.Error()}})
	}
	var results []models.Camp
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
	location, err := utills.DecodeUrl(location)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "location is invalid",
			Data:    &fiber.Map{"data": err.Error()}})
	}

	filter := bson.D{{}}
	campsData, err := campCollection.Find(context.TODO(), filter)

	var results []models.Camp

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{
			Status:  http.StatusNotFound,
			Message: "camp in this keyword not found",
			Data:    &fiber.Map{"data": err.Error()}})
	}

	if err = campsData.All(context.TODO(), &results); err != nil {
		return c.Status(http.StatusBadGateway).JSON(responses.UserResponse{Status: http.StatusBadGateway, Message: "error", Data: &fiber.Map{"error": err.Error()}})
	}
	var queryCampsData []models.Camp

	for i := 0; i < len(results); i++ {
		if utills.MatchString(results[i].Location, location) {
			queryCampsData = append(queryCampsData, results[i])
		}
	}
	return c.Status(http.StatusAccepted).JSON(responses.UserResponse{Status: http.StatusAccepted, Message: "success", Data: &fiber.Map{"camps": queryCampsData}})
}

func GetCampDatabyKeyword(c *fiber.Ctx) error {
	keyword := c.Params("keyword")
	keyword, err := utills.DecodeUrl(keyword)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "keyword is invalid",
			Data:    &fiber.Map{"data": err.Error()}})
	}

	filter := bson.D{{}}
	campsData, err := campCollection.Find(context.TODO(), filter)

	var results []models.Camp

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{
			Status:  http.StatusNotFound,
			Message: "camp with this keyword not found",
			Data:    &fiber.Map{"data": err.Error()}})
	}

	if err = campsData.All(context.TODO(), &results); err != nil {
		return c.Status(http.StatusBadGateway).JSON(responses.UserResponse{Status: http.StatusBadGateway, Message: "error", Data: &fiber.Map{"error": err.Error()}})
	}
	var queryCampsData []models.Camp

	for i := 0; i < len(results); i++ {
		var matched bool = utills.MatchString(results[i].Location, keyword) || utills.MatchString(results[i].Name, keyword) || utills.MatchString(results[i].Director, keyword)
		if matched {
			queryCampsData = append(queryCampsData, results[i])
		}
	}
	return c.Status(http.StatusAccepted).JSON(responses.UserResponse{Status: http.StatusAccepted, Message: "success", Data: &fiber.Map{"camps": queryCampsData}})
}

func GetCampDatabyYear(c *fiber.Ctx) error {
	year := c.Params("year")
	year, err := utills.DecodeUrl(year)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "year is invalid",
			Data:    &fiber.Map{"data": err.Error()}})
	}

	filter := bson.D{{}}
	campsData, err := campCollection.Find(context.TODO(), filter)

	var results []models.Camp

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{
			Status:  http.StatusNotFound,
			Message: "camp in this year not found",
			Data:    &fiber.Map{"data": err.Error()}})
	}

	if err = campsData.All(context.TODO(), &results); err != nil {
		return c.Status(http.StatusBadGateway).JSON(responses.UserResponse{Status: http.StatusBadGateway, Message: "error", Data: &fiber.Map{"error": err.Error()}})
	}
	var queryCampsData []models.Camp

	for i := 0; i < len(results); i++ {
		var camptime []string = strings.Split(results[i].Time, " ")
		if len(camptime) > 1 {
			var campyear string = camptime[1]
			if campyear == year {
				queryCampsData = append(queryCampsData, results[i])
			}
		}
	}
	return c.Status(http.StatusAccepted).JSON(responses.UserResponse{Status: http.StatusAccepted, Message: "success", Data: &fiber.Map{"camps": queryCampsData}})
}
