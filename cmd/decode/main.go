package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/rafaeltorres324/multicode/cmd/decode/database"
)

var verbose bool
var AllData []TokenData
var CharlesData CharlesLog
var SentData int
var root = "/Users/support1/CharlesFNLExports"

func init() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Init()
}

func main() {

	for {
		CheckFiles()
		log.Println("Waiting for more files..")
		time.Sleep(20 * time.Second)
	}
}

func CheckFiles() {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	for _, file := range files {
		if !strings.Contains(file, "chlsj") || strings.Contains(file, "used") {
			continue
		}
		fmt.Println(file)
		UseFile(file)
	}
}

func UseFile(file string) {
	jsonData, err := os.Open(file)
	if err != nil {
		log.Println(err)
	}

	dataFile, _ := ioutil.ReadAll(jsonData)

	json.Unmarshal(dataFile, &CharlesData)

	log.Println(len(CharlesData))
	jsonData.Close()

	for f := range CharlesData {
		if CharlesData[f].Host == "www.recaptcha.net" {
			if strings.Contains(CharlesData[f].Path, "iosc") {
				continue
			}
			StartDecoding([]byte(CharlesData[f].Request.Body.Encoded))
		}
	}

	MoveFile(file)

	log.Println(len(AllData))

	SendData()

	log.Println("Total data sent", SentData)
	AllData = nil
}

func MoveFile(file string) {
	Original_Path := file
	New_Path := "/Users/support1/CharlesFNLExports/used/" + GetTimestamp() + ".chlsj"
	log.Println("Renaming file", Original_Path, "to", New_Path)
	time.Sleep(2 * time.Second)
	e := os.Rename(Original_Path, New_Path)
	if e != nil {
		log.Println(e)
	}
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
