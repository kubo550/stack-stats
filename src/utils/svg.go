package utils

import (
	"stats/src/structs"
)

func GenerateSVG(stackStats structs.Stats, theme structs.Theme) (string, error) {
	const width = 158
	const height = 47
	const fontSize = 12
	const badgeStartPosPx = 90
	const badgesGap = 22

	var svg string

	svg += `<svg data-testUserId="` + stackStats.ID + `" width="` + str(width) + `" height="` + str(height) + `" viewBox="0 0 ` + str(width) + ` ` + str(height) + `" fill="none" xmlns="http://www.w3.org/2000/svg">`
	svg += `<rect width="` + str(width) + `" height="` + str(height) + `" fill="` + theme.BgColor + `"/>`

	// Profile image

	if stackStats.ImageUrl != "" {
		imageBase64, err := ImageToBase64(stackStats.ImageUrl)
		if err != nil {
			return "", err
		}
		svg += generateImage(imageBase64)
	}
	// Reputation
	svg += displayReputation(stackStats, theme, height, fontSize)

	// Gold
	svg += generateBadge("Gold", badgeStartPosPx, height/2, stackStats.Gold, fontSize, theme.Gold)

	// Silver
	svg += generateBadge("Silver", badgeStartPosPx+badgesGap, height/2, stackStats.Silver, fontSize, theme.Silver)

	// Bronze
	svg += generateBadge("Bronze", badgeStartPosPx+2*badgesGap, height/2, stackStats.Bronze, fontSize, theme.Bronze)

	svg += `</svg>`

	return svg, nil
}

func displayReputation(stackStats structs.Stats, theme structs.Theme, height int, fontSize int) string {
	svg := `<text data-testReputation="` + formatNumberWithComma(stackStats.Reputation) + `"  x="` + str(64) + `" y="` + str(height/2) + `" font-weight="bold" fill="` + theme.TextColor + `" font-family="Arial" font-size="` + str(fontSize) + `" text-anchor="middle" dominant-baseline="middle">` + formatNumberWithComma(stackStats.Reputation) + `</text>`
	return svg
}

func generateImage(imageBase64 string) (svg string) {
	fullImage := "data:image/png;base64," + imageBase64
	svg = ` <image data-testImageUrl="` + fullImage + `" x="16" y="10" href="` + fullImage + `" height="24" width="24"/>`

	return svg
}

func generateBadge(id string, xPos, yPos, count, fontSize int, color string) (svg string) {
	if count == 0 {
		return svg
	}

	gap := calculateGap(count)
	const radius = 3

	svg += `<circle text-anchor="middle" dominant-baseline="middle" cx="` + str(xPos) + `" cy="` + str(yPos) + `" r="` + str(radius) + `" fill="` + color + `"/>`
	svg += `<text data-testBadge` + id + `="` + formatNumberWithComma(count) + `" x="` + str(xPos+gap) + `" y="` + str(yPos) + `" font-size="` + str(fontSize) + `" font-family="Arial" font-weight="bold" text-anchor="middle" dominant-baseline="middle" fill="` + color + `">` + formatNumberWithComma(count) + `</text>`

	return svg

}

func calculateGap(count int) int {
	var gap = 10
	if count > 10 {
		gap = 12
	} else if count > 100 {
		gap = 14
	} else if count > 1000 {
		gap = 16
	}
	return gap
}
