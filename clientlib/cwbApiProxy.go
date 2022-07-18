package clientlib

import (
	"JimLineBot-v2/definition"
	"encoding/json"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func GetRadarPicUri() (string, error) {
	godotenv.Load()
	t := time.Now()
	cwbAuthorizationToken := os.Getenv("CWB_API_AUTHORIZATION")
	resp, err := http.Get(definition.RadarRequest + cwbAuthorizationToken + "&format=JSON")
	if err != nil {
		return "", err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var jsonObj map[string]interface{}
	json.Unmarshal(body, &jsonObj)
	cwbopendata := jsonObj["cwbopendata"].(map[string]interface{})
	dataset := cwbopendata["dataset"].(map[string]interface{})
	resource := dataset["resource"].(map[string]interface{})
	uri := resource["uri"].(string)
	uri = uri + "?" + strconv.Itoa(t.Day()) + strconv.Itoa(t.Hour()) + strconv.Itoa(t.Minute()) + strconv.Itoa(t.Second())
	defer resp.Body.Close()
	return uri, nil
}
