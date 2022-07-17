package clientlib

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func lineNotify(uri string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://notify-api.line.me/api/notify", nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Printf("%s", result)
	return result, nil
}
