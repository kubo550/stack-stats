package utils

import (
	"stats/src/structs"
)

func GenerateSVG(stackStats structs.Stats, theme structs.Theme) (svg string, error error) {
	const height = 47
	const fontSize = 12
	const badgesGap = 18
	const outerPadding = 12
	const innerPadding = 6
	const imageSize = 24

	width := calcWidth(stackStats.Reputation, stackStats.Gold, stackStats.Silver, stackStats.Bronze, badgesGap)

	svg += `<svg data-testUserId="` + stackStats.ID + `" width="` + str(width) + `" height="` + str(height) + `" viewBox="0 0 ` + str(width) + ` ` + str(height) + `" fill="none" xmlns="http://www.w3.org/2000/svg">`
	svg += `<rect width="` + str(width) + `" height="` + str(height) + `" fill="` + theme.BgColor + `"/>`

	if stackStats.ImageUrl != "" {
		imageBase64, err := ImageToBase64(stackStats.ImageUrl)
		if err != nil {
			return "", err
		}
		svg += generateImage(imageBase64, outerPadding, imageSize)
	}

	svg += displayReputation(outerPadding+imageSize+innerPadding+18, stackStats, theme, height, fontSize)

	badgeXPos := outerPadding + imageSize + innerPadding*2 + calcReputationWidth(stackStats.Reputation) + innerPadding*2
	if stackStats.Gold > 0 {
		svg += generateBadge("Gold", badgeXPos, height/2, stackStats.Gold, fontSize, theme.Gold)
	}

	if stackStats.Silver > 0 {
		badgeXPos = badgeXPos + calcBadgeScoreGap(stackStats.Gold) + badgesGap
		svg += generateBadge("Silver", badgeXPos, height/2, stackStats.Silver, fontSize, theme.Silver)
	}

	if stackStats.Bronze > 0 {
		badgeXPos = badgeXPos + calcBadgeScoreGap(stackStats.Silver) + badgesGap
		svg += generateBadge("Bronze", badgeXPos, height/2, stackStats.Bronze, fontSize, theme.Bronze)
	}

	svg += `</svg>`

	return svg, nil
}

func calcWidth(reputation, gold, silver, bronze, badgesGap int) int {
	minWidth := 83
	scaler := 0.5
	for _, b := range [...]int{gold, silver, bronze} {
		if b > 0 {

			minWidth += calcBadgeScoreGap(b) + scale(calcBadgeScoreGap(b), scaler) + scale(badgesGap, scaler)
		}
	}
	width := minWidth + calcReputationWidth(reputation)

	return width

}

func scale(value int, factor float64) int {
	return int(float64(value) * factor)
}

func displayReputation(xPos int, stackStats structs.Stats, theme structs.Theme, height int, fontSize int) string {
	svg := `<text data-testReputation="` + formatNumberWithComma(stackStats.Reputation) + `"  x="` + str(xPos) + `" y="` + str(height/2) + `" font-weight="bold" fill="` + theme.TextColor + `" font-family="Arial" font-size="` + str(fontSize) + `" text-anchor="middle" dominant-baseline="middle">` + formatNumberWithComma(stackStats.Reputation) + `</text>`
	return svg
}

func generateImage(imageBase64 string, xPos, size int) (svg string) {
	fullImage := "data:image/png;base64," + imageBase64
	svg = ` <image x=" ` + str(xPos) + `" y="10" href="` + fullImage + `" height="` + str(size) + `" width="` + str(size) + `" default-src="sha256-4Su6mBWzEIFnH4pAGMOuaeBrstwJN4Z3pq/s1Kn4/KQ=" />`

	return svg
}

func generateBadge(id string, xPos, yPos, count, fontSize int, color string) (svg string) {

	gap := calcBadgeScoreGap(count)
	const radius = 3

	svg += `<circle text-anchor="middle" dominant-baseline="middle" cx="` + str(xPos) + `" cy="` + str(yPos) + `" r="` + str(radius) + `" fill="` + color + `"/>`
	svg += `<text data-testBadge` + id + `="` + formatNumberWithComma(count) + `" x="` + str(xPos+gap) + `" y="` + str(yPos) + `" font-size="` + str(fontSize) + `" font-family="Arial" font-weight="bold" text-anchor="middle" dominant-baseline="middle" fill="` + color + `">` + formatNumberWithComma(count) + `</text>`

	return svg

}

func calcBadgeScoreGap(count int) int {
	if count == 0 {
		return 0
	} else if count < 10 {
		return 9
	} else if count < 100 {
		return 14
	} else if count < 1000 {
		return 16
	}
	return 18
}
func calcReputationWidth(reputation int) int {
	if reputation < 10 {
		return 7
	} else if reputation < 100 {
		return 12
	} else if reputation < 1000 {
		return 20
	}
	return 33
}
