package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Monitor struct {
	Id                int
	TestId            string `json:"testId"`
	JsonUrl           string `json:"jsonUrl"`
	WebTestingUrlConf string `json:"webTestingUrlConf"`
	SummaryUrl        string `json:"summaryUrl"`
	Har               string `json:"har"`
	Platform          string `json:"platform"`
	PublisherId       string `json:"publisherId"`
	CreatedAt         string `json:"createdAt"`
	IsError           string `json:"isError"`
}

var (
	id                int
	testId            string
	jsonUrl           string
	webTestingUrlConf string
	summaryUrl        string
	har               string
	platform          string
	publisher_id      string
	createdAt         string
	isError           string
)

func GetSearchResult(resultStr string) []Monitor {
	var arrMain []Monitor
	//DB connection
	db, err := sql.Open("mysql", "<db-info>")
	if err != nil {
		log.Fatal("Error DB connection: ", err.Error())
	}
	defer db.Close()

	/**
	 * Preparing statement for reading DB
	 */
	rows, err := db.Query("SELECT id, testId, jsonUrl, webTestingUrlConf, summaryUrl, har, platform, publisher_id, createdAt, isError FROM analysis_tool_trc.logger_web_page WHERE publisher_id LIKE ? ORDER BY createdAt DESC LIMIT 300", "%"+resultStr+"%")
	if err != nil {
		log.Fatal("Error while fetching results: ", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &testId, &jsonUrl, &webTestingUrlConf, &summaryUrl, &har, &platform, &publisher_id, &createdAt, &isError)
		if err != nil {
			log.Fatal("Error rows in loop: ", err)
		}

		monitorObj := Monitor{id, testId, jsonUrl, webTestingUrlConf, summaryUrl, har, platform, publisher_id, createdAt, isError}

		jsonObj := Monitor{
			Id:                monitorObj.Id,
			TestId:            monitorObj.TestId,
			JsonUrl:           monitorObj.JsonUrl,
			WebTestingUrlConf: monitorObj.WebTestingUrlConf,
			SummaryUrl:        monitorObj.SummaryUrl,
			Har:               monitorObj.Har,
			Platform:          monitorObj.Platform,
			PublisherId:       monitorObj.PublisherId,
			CreatedAt:         monitorObj.CreatedAt,
			IsError:           monitorObj.IsError,
		}
		arrMain = append(arrMain, jsonObj)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal("Error in rows: ", err)
	}

	return arrMain
}
