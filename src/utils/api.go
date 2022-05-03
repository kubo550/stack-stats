package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"stats/src/structs"
)

func GetStackStats(userId string) structs.Stats {
	client := &http.Client{}
	req, err := http.NewRequest(fiber.MethodGet, getStackApiUrl(userId), nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if err != nil {
		panic(err)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var stackStats structs.StackStats
	err = json.Unmarshal(responseBody, &stackStats)
	if err != nil {
		log.Fatal(err)
	}

	return structs.Stats{
		ID:         userId,
		Name:       stackStats.Items[0].DisplayName,
		Reputation: stackStats.Items[0].Reputation,
		Gold:       stackStats.Items[0].BadgeCounts.Gold,
		Silver:     stackStats.Items[0].BadgeCounts.Silver,
		Bronze:     stackStats.Items[0].BadgeCounts.Bronze,
		ImageUrl:   stackStats.Items[0].ProfileImage,
	}
}

func getStackApiUrl(userId string) string {
	return fmt.Sprintf("https://api.stackexchange.com/2.3/users/%s?order=desc&sort=reputation&site=stackoverflow", userId)
}
