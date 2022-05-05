package builders

import (
	"gopkg.in/h2non/gock.v1"
	"stats/src/structs"
)

func StackExchangeWillRespondWith(status int, response structs.StackResponse) {
	gock.New("https://api.stackexchange.com").
		Get("/2.3/users/*").
		Reply(status).
		JSON(response)
}

type StackResponseBuilder interface {
	WithName(name string) StackResponseBuilder
	WithReputation(reputation int) StackResponseBuilder
	WithBadgeCounts(badgeCounts structs.BadgeCounts) StackResponseBuilder
	WithImageUrl(imgUrl string) StackResponseBuilder
	Build() structs.StackResponse
}

func NewStackResponseBuilder() StackResponseBuilder {
	// default values
	return &stackResponseBuilderImpl{
		name:       "John Doe",
		reputation: 50,
		badgeCounts: structs.BadgeCounts{
			Bronze: 1,
			Silver: 2,
			Gold:   3,
		},
		imgUrl: "https:image.com",
	}
}

type stackResponseBuilderImpl struct {
	name        string
	reputation  int
	badgeCounts structs.BadgeCounts
	imgUrl      string
}

func (b *stackResponseBuilderImpl) WithName(name string) StackResponseBuilder {
	b.name = name
	return b
}

func (b *stackResponseBuilderImpl) WithReputation(reputation int) StackResponseBuilder {
	b.reputation = reputation
	return b
}

func (b *stackResponseBuilderImpl) WithBadgeCounts(badgeCounts structs.BadgeCounts) StackResponseBuilder {
	b.badgeCounts = badgeCounts
	return b
}

func (b *stackResponseBuilderImpl) WithImageUrl(imgUrl string) StackResponseBuilder {
	b.imgUrl = imgUrl
	return b
}

func (b *stackResponseBuilderImpl) Build() structs.StackResponse {
	return generateStackResponse(b.name, b.reputation, b.badgeCounts, b.imgUrl)
}

func generateStackResponse(name string, reputation int, badgeCounts structs.BadgeCounts, imgUrl string) structs.StackResponse {
	return structs.StackResponse{
		Items: []struct {
			BadgeCounts struct {
				Bronze int `json:"bronze"`
				Silver int `json:"silver"`
				Gold   int `json:"gold"`
			} `json:"badge_counts"`
			Reputation   int    `json:"reputation"`
			ProfileImage string `json:"profile_image"`
			DisplayName  string `json:"display_name"`
		}{
			{
				BadgeCounts:  badgeCounts,
				Reputation:   reputation,
				ProfileImage: imgUrl,
				DisplayName:  name,
			},
		},
	}
}
