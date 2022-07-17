package clientlib

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
)

func GetRadarPicUri() (string, error) {
	godotenv.Load()
	cwbAuthorizationToken := os.Getenv("CWB_API_AUTHORIZATION")
	resp, err := http.Get("definition.RadarRequest" + cwbAuthorizationToken + "&format=JSON")
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
	defer resp.Body.Close()
	return uri, nil
}
