package utils

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"net/http"
	"stats/src/log"
)

func ImageToBase64(imageUrl string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", imageUrl, nil)
	if err != nil {
		return "", err
	}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(err)
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	base64Data := base64.StdEncoding.EncodeToString(body)
	return base64Data, nil
}
