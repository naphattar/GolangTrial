package scrappers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/naphattar/KaihorBackend/configs"
	"github.com/naphattar/KaihorBackend/models"
	"github.com/naphattar/KaihorBackend/responses"
)

func printCampRawData(campData responses.SheetapiResponse) {
	fmt.Println("--------------------------------")
	fmt.Println("id : ", campData.CampID)
	fmt.Println("time : ", campData.Time)
	fmt.Println("name : ", campData.Name)
	fmt.Println("location : ", campData.Location)
	fmt.Println("director : ", campData.Director)
}

func validateCampJSON(campJSON map[string]string) models.Camp {
	campid := campJSON["ค่ายรวมจุฬาฯ อาสาพัฒนาชนบท โดยชมรมค่ายอาสาสมัครหอพักนิสิตจุฬาลงกรณ์มหาวิทยาลัย"]
	campname := campJSON["3"]
	camptime := campJSON["1"]
	camplocation := campJSON["2"]
	campdirector := campJSON["4"]

	if campid == "" {
		campid = "ไม่ระบุ"
	}

	if campname == "" {
		campname = "ไม่ระบุ"
	}
	if camptime == "" {
		camptime = "ไม่ระบุ"
	}
	if camplocation == "" {
		camplocation = "ไม่ระบุ"
	}
	if campdirector == "" {
		campdirector = "ไม่ระบุ"
	}
	newCamp := models.Camp{
		CampID:   campid,
		Name:     campname,
		Time:     camptime,
		Location: camplocation,
		Director: campdirector,
	}
	return newCamp
}

func GetDatafromSheet() []models.Camp {
	sheetapi := configs.EnvSpreadSheetAPI()
	response, err := http.Get(sheetapi)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// turn raw data to map (array of JSON)
	var responseDataJSON []map[string]string

	if err := json.Unmarshal(responseData, &responseDataJSON); err != nil {
		log.Fatal(err)
	}

	// array of Camps to store all datas
	var campsData []models.Camp

	for i := 0; i < len(responseDataJSON); i++ {
		newCamp := validateCampJSON(responseDataJSON[i])
		campsData = append(campsData, newCamp)
	}
	return campsData
}
