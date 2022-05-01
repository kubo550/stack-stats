package utils

import (
	"stats/src/structs"
)

func GenerateSVG(stackStats structs.Stats, theme structs.Theme) string {
	const width = 158
	const height = 47
	const fontSize = 12
	const badgeStartPosPx = 90
	const badgesGap = 22

	var svg string

	svg += `<svg width="` + str(width) + `" height="` + str(height) + `" viewBox="0 0 ` + str(width) + ` ` + str(height) + `" fill="none" xmlns="http://www.w3.org/2000/svg">`
	svg += `<rect width="` + str(width) + `" height="` + str(height) + `" fill="#2D2D2D"/>`

	// Profile image
	if stackStats.ImageUrl != "" {
		svg += ` <image x="16" y="10" href="` + stackStats.ImageUrl + `" height="24" width="24"/>`
	}

	// Reputation
	svg += `<text x="` + str(64) + `" y="` + str(height/2) + `" font-weight="bold" fill="#C4CCBC" font-family="Arial" font-size="` + str(fontSize) + `" text-anchor="middle" dominant-baseline="middle">` + formatNumber(stackStats.Reputation) + `</text>`

	// Gold
	svg += generateBadge(badgeStartPosPx, height/2, stackStats.Gold, fontSize, theme.Gold)

	// Silver
	svg += generateBadge(badgeStartPosPx+badgesGap, height/2, stackStats.Silver, fontSize, theme.Silver)

	// Bronze
	svg += generateBadge(badgeStartPosPx+2*badgesGap, height/2, stackStats.Bronze, fontSize, theme.Bronze)

	svg += `</svg>`

	return svg
}

func generateBadge(xPos, yPos, count, fontSize int, color string) string {
	svg := ""
	if count == 0 {
		return svg
	}

	gap := calculateGap(count)
	const radius = 3

	svg += `<circle text-anchor="middle" dominant-baseline="middle" cx="` + str(xPos) + `" cy="` + str(yPos) + `" r="` + str(radius) + `" fill="` + color + `"/>`
	svg += `<text x="` + str(xPos+gap) + `" y="` + str(yPos) + `" font-size="` + str(fontSize) + `" font-family="Arial" font-weight="bold" text-anchor="middle" dominant-baseline="middle" fill="` + color + `">` + formatNumber(count) + `</text>`

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