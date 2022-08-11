package main

import (
	"JimLineBot-v2/clientlib"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"net/http"
	"os"
)

var (
	client *linebot.Client
	err    error
)

func main() {
	godotenv.Load()
	// 建立客戶端
	fmt.Println("channelSecret:" + os.Getenv("CHANNEL_SECRET"))
	//fmt.Println(os.Getenv("CHANNEL_ACCESS_TOKEN"))
	client, err = linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_ACCESS_TOKEN"))

	if err != nil {
		log.Println(err.Error())
	}
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	// 接收請求
	events, err := client.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}

		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				// 回覆訊息
				if message.Text == ("雷達回波") || message.Text == ("雷達會破") || message.Text == ("雷達迴波") {
					uri, err := clientlib.GetRadarPicUri()
					if err != nil {
						return
					}
					_, err = client.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(uri, uri)).Do()
					if err != nil {
						log.Println(err.Error())
					}
				}
				if message.Text == ("大學") {
					var albumId = "JRRhuxN"
					uri, err := clientlib.GetImageUriFromimgur(albumId)
					if err != nil {
						return
					}
					_, err = client.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(uri, uri)).Do()
					if err != nil {
						log.Println(err.Error())
					}
				}
				if _, err = client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Println(err.Error())
				}
			}
		}
	}
}
