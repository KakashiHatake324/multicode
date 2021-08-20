package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func SendData() {
	client := http.Client{}

	for i := range AllData {

		json, _ := json.Marshal(AllData[i])

		body := bytes.NewReader(json)

		req, err := http.NewRequest("POST", "http://akamai.kageaio.com:7821/PostData", body)
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
			}

			log.Println(string(bodyText))
			defer resp.Body.Close()

		}

	}
}
