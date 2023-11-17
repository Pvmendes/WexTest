package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 60 * time.Second}

func GetJson(urlGet string, target interface{}) error {
	r, err := myClient.Get(urlGet)

	if err != nil {
		return err
	}

	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}