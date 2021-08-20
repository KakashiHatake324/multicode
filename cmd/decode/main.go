package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var verbose bool
var AllData []TokenData
var CharlesData CharlesLog

func main() {

	jsonData, err := os.Open("/Users/support1/CharlesFNLExports/Untitled.chlsj")
	if err != nil {
		log.Println(err)
	}

	dataFile, _ := ioutil.ReadAll(jsonData)

	json.Unmarshal(dataFile, &CharlesData)

	log.Println(CharlesData[1].Host)
	log.Println(len(CharlesData))
	jsonData.Close()

	for f := range CharlesData {

		if CharlesData[f].Host != "www.recaptcha.net" {
			continue
		}
		StartDecoding([]byte(CharlesData[f].Request.Body.Encoded))

	}

	log.Println(len(AllData))

	SendData()
}

func StartDecoding(input []byte) {

	data := GetValues(input)

	text := string(data)
	text = strings.ReplaceAll(text, "1:", "")
	text = strings.ReplaceAll(text, "2:", "")
	text = strings.ReplaceAll(text, "3:", "")
	text = strings.ReplaceAll(text, "4:", "")
	splitData := strings.Split(text, " ")

	var NewToken TokenData
	NewToken.Token = strings.ReplaceAll(splitData[0], "\"", "")
	NewToken.Action = strings.ReplaceAll(splitData[1], "\"", "")
	NewToken.Timestamp, _ = strconv.ParseInt(splitData[2], 10, 64)
	NewToken.DeviceInfo = strings.ReplaceAll(splitData[3], "\"", "")

	AllData = append(AllData, NewToken)

}

type TokenData struct {
	Token      string `json:"token"`
	Action     string `json:"action"`
	Timestamp  int64  `json:"timestamp"`
	DeviceInfo string `json:"deviceInfo"`
}

type ReCaptchaRequest struct {
	Token      string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Action     string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Timestamp  int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	DeviceInfo string `protobuf:"bytes,4,opt,name=device_info,json=deviceInfo,proto3" json:"device_info,omitempty"`
}

type CharlesLog []struct {
	Status          string      `json:"status"`
	Method          string      `json:"method"`
	ProtocolVersion string      `json:"protocolVersion"`
	Scheme          string      `json:"scheme"`
	Host            string      `json:"host"`
	ActualPort      int         `json:"actualPort"`
	Path            string      `json:"path"`
	Query           interface{} `json:"query"`
	Tunnel          bool        `json:"tunnel"`
	KeptAlive       bool        `json:"keptAlive"`
	WebSocket       bool        `json:"webSocket"`
	RemoteAddress   string      `json:"remoteAddress"`
	ClientAddress   string      `json:"clientAddress"`
	ClientPort      int         `json:"clientPort"`
	Times           struct {
		Start           string `json:"start"`
		RequestBegin    string `json:"requestBegin"`
		RequestComplete string `json:"requestComplete"`
		ResponseBegin   string `json:"responseBegin"`
		End             string `json:"end"`
	} `json:"times"`
	Durations struct {
		Total    int `json:"total"`
		DNS      int `json:"dns"`
		Connect  int `json:"connect"`
		Ssl      int `json:"ssl"`
		Request  int `json:"request"`
		Response int `json:"response"`
		Latency  int `json:"latency"`
	} `json:"durations"`
	Speeds struct {
		Overall  int `json:"overall"`
		Request  int `json:"request"`
		Response int `json:"response"`
	} `json:"speeds"`
	TotalSize int `json:"totalSize"`
	Ssl       struct {
		Protocol    string `json:"protocol"`
		CipherSuite string `json:"cipherSuite"`
	} `json:"ssl"`
	Alpn struct {
		Protocol string `json:"protocol"`
	} `json:"alpn"`
	Request struct {
		Sizes struct {
			Handshake int `json:"handshake"`
			Headers   int `json:"headers"`
			Body      int `json:"body"`
		} `json:"sizes"`
		MimeType        string      `json:"mimeType"`
		Charset         interface{} `json:"charset"`
		ContentEncoding interface{} `json:"contentEncoding"`
		Header          struct {
			Headers []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"headers"`
		} `json:"header"`
		Body struct {
			Encoding string `json:"encoding"`
			Encoded  string `json:"encoded"`
		} `json:"body"`
	} `json:"request"`
	Response struct {
		Status int `json:"status"`
		Sizes  struct {
			Handshake int `json:"handshake"`
			Headers   int `json:"headers"`
			Body      int `json:"body"`
		} `json:"sizes"`
		MimeType        string      `json:"mimeType"`
		Charset         interface{} `json:"charset"`
		ContentEncoding string      `json:"contentEncoding"`
		Header          struct {
			Headers []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"headers"`
		} `json:"header"`
		Body struct {
			Text    string      `json:"text"`
			Charset interface{} `json:"charset"`
			Decoded bool        `json:"decoded"`
		} `json:"body"`
	} `json:"response,omitempty"`
}

type Request struct {
	Sizes struct {
		Handshake int `json:"handshake"`
		Headers   int `json:"headers"`
		Body      int `json:"body"`
	} `json:"sizes"`
	MimeType        interface{} `json:"mimeType"`
	Charset         interface{} `json:"charset"`
	ContentEncoding interface{} `json:"contentEncoding"`
	Header          struct {
		FirstLine string `json:"firstLine"`
		Headers   []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"headers"`
	} `json:"header"`
}

type Response struct {
	Status int `json:"status"`
	Sizes  struct {
		Handshake int `json:"handshake"`
		Headers   int `json:"headers"`
		Body      int `json:"body"`
	} `json:"sizes"`
	MimeType        interface{} `json:"mimeType"`
	Charset         interface{} `json:"charset"`
	ContentEncoding interface{} `json:"contentEncoding"`
	Header          struct {
		FirstLine string `json:"firstLine"`
		Headers   []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"headers"`
	} `json:"header"`
	Body struct {
		Text    string      `json:"text"`
		Charset interface{} `json:"charset"`
	} `json:"body"`
}
