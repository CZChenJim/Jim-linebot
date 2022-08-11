package clientlib

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func GetImageUriFromimgur(albumId string) (string, error) {
	godotenv.Load()
	imgurClientId := os.Getenv("IMGUR_CLIENT_ID")
	req, err := http.NewRequest("GET", "https://api.imgur.com/3/album/"+imgurClientId+"/images", nil)
	req.Header.Set("Authorization", "Client-ID "+imgurClientId)
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return "", err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var jsonObj map[string]interface{}
	json.Unmarshal(body, &jsonObj)
	dataList := jsonObj["data"].([]interface{})
	rand.Seed(time.Now().Unix())
	data := dataList[rand.Intn(len(dataList))].(map[string]interface{})
	link := data["link"].(string)
	defer resp.Body.Close()
	return link, nil
}
