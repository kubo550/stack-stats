package structs

type Theme struct {
	Gold      string `json:"gold"`
	Silver    string `json:"silver"`
	Bronze    string `json:"bronze"`
	BgColor   string `json:"bg_color"`
	TextColor string `json:"text_color"`
}

type Stats struct {
	ID         string `json:"id"`
	Name       string `json:"fullName"`
	Reputation int    `json:"reputation"`
	Gold       int    `json:"gold"`
	Silver     int    `json:"silver"`
	Bronze     int    `json:"bronze"`
	ImageUrl   string `json:"imageUrl"`
}

type StackStats struct {
	Items []struct {
		DisplayName  string `json:"display_name"`
		ProfileImage string `json:"profile_image"`
		Reputation   int    `json:"reputation"`
		BadgeCounts  struct {
			Bronze int `json:"bronze"`
			Gold   int `json:"gold"`
			Silver int `json:"silver"`
		} `json:"badge_counts"`
	} `json:"items"`
}

type BadgeCounts struct {
	Bronze int `json:"bronze"`
	Silver int `json:"silver"`
	Gold   int `json:"gold"`
}

type StackResponse struct {
	Items []struct {
		BadgeCounts struct {
			Bronze int `json:"bronze"`
			Silver int `json:"silver"`
			Gold   int `json:"gold"`
		} `json:"badge_counts"`
		Reputation   int    `json:"reputation"`
		ProfileImage string `json:"profile_image"`
		DisplayName  string `json:"display_name"`
	} `json:"items"`
}

type StackResponseBuilder interface {
	WithName(name string) StackResponseBuilder
	WithReputation(reputation int) StackResponseBuilder
	WithBadgeCounts(badgeCounts BadgeCounts) StackResponseBuilder
	WithImageUrl(imgUrl string) StackResponseBuilder
	Build() StackResponse
}
