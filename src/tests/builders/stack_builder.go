package builders

import (
	"stats/src/structs"
)

func NewStackResponseBuilder() structs.StackResponseBuilder {
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

func (b *stackResponseBuilderImpl) WithName(name string) structs.StackResponseBuilder {
	b.name = name
	return b
}

func (b *stackResponseBuilderImpl) WithReputation(reputation int) structs.StackResponseBuilder {
	b.reputation = reputation
	return b
}

func (b *stackResponseBuilderImpl) WithBadgeCounts(badgeCounts structs.BadgeCounts) structs.StackResponseBuilder {
	b.badgeCounts = badgeCounts
	return b
}

func (b *stackResponseBuilderImpl) WithImageUrl(imgUrl string) structs.StackResponseBuilder {
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
