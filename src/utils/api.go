package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"io/ioutil"
	"net/http"
	"stats/src/log"
	"stats/src/structs"
)

func GetStackStats(userId string) (structs.Stats, error) {
	client := &http.Client{}
	req, err := http.NewRequest(fiber.MethodGet, getStackApiUrl(userId), nil)
	if err != nil {
		return structs.Stats{}, err
	}

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != fiber.StatusOK {
		return structs.Stats{}, fmt.Errorf("error getting stack stats")
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(err)
		}
	}(resp.Body)

	if err != nil {
		return structs.Stats{}, err
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return structs.Stats{}, err
	}

	var stackStats structs.StackStats

	err = json.Unmarshal(responseBody, &stackStats)

	if err != nil {
		log.Error(err)
		return structs.Stats{}, err
	}

	return structs.Stats{
		ID:         userId,
		Name:       stackStats.Items[0].DisplayName,
		Reputation: stackStats.Items[0].Reputation,
		Gold:       stackStats.Items[0].BadgeCounts.Gold,
		Silver:     stackStats.Items[0].BadgeCounts.Silver,
		Bronze:     stackStats.Items[0].BadgeCounts.Bronze,
		ImageUrl:   stackStats.Items[0].ProfileImage,
	}, nil
}

func getStackApiUrl(userId string) string {
	return fmt.Sprintf("https://api.stackexchange.com/2.3/users/%s?order=desc&sort=reputation&site=stackoverflow", userId)
}
