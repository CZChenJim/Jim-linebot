package clientlib

import (
	"JimLineBot-v2/definition"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func LineNotify(uri string) ([]byte, error) {
	godotenv.Load()
	notifyToken := os.Getenv("LINE_NOTIFY_TOKEN")
	client := &http.Client{}
	data := make(map[string]string)
	data["message"] = "從雷達回波看看會不會下雨～"
	data["imageThumbnail"] = uri
	data["imageFullsize"] = uri
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(data)
	fmt.Println("try2")
	req, err := http.NewRequest("POST", definition.NotifyRequest, b)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+notifyToken)
	fmt.Println("try3")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Println("try")
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println("%s", result)
	return result, nil
}
