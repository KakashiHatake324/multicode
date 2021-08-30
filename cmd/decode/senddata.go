package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/rafaeltorres324/multicode/cmd/decode/database"
)

type CollectionsData struct {
	Payload string `json:"payload"`
}

func SendData() {
	//client := http.Client{}

	collection := database.Instance.Client.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("TOKENS_COLLECTION"))

	for i := range AllData {

		// Json marshal the data
		json, _ := json.Marshal(AllData[i])
		// Encode the data to hex in order to save space
		hx := hex.EncodeToString([]byte(string(json)))

		dbBody := CollectionsData{
			Payload: hx,
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

		//log.Println(string(bs))

		/*
			body := bytes.NewReader(json)

			req, err := http.NewRequest("POST", server, body)
			if err != nil {
				log.Println("Error sending data..")
			}
			req.Header.Set("Content-Type", "Application/Json")
			resp, err := client.Do(req)
			if err != nil {
				log.Println("Error sending data..")
			} else {

				bodyText, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Println("Error sending data..")
				} else {
					SentData++
				}

				log.Println(string(bodyText))
				defer resp.Body.Close()

			}
		*/
	}
}
