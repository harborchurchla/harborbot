package utils

import (
	"fmt"
	"strings"
)

func ReplaceNameWithSlackMention(text string) string {
	slackUserMap := map[string]string{
		"peppy":  "U010DDYSGDN",
		"shalom": "U0106QYAU03",
		"ida":    "U010DGBGMUG",
		"angelo": "U010HN4HH60",
		"angel":  "U010HN4HH60",
		"bikki":  "U010HUMG7TQ",
		"caleb":  "U0105LQ01A5",
		"jackie": "U010KU8141M",
		"benu":   "U01JTK7SEM9",
	}

	for name, slackUserId := range slackUserMap {
		text = strings.ReplaceAll(text, name, fmt.Sprintf("<@%s>", slackUserId))
		text = strings.ReplaceAll(text, strings.Title(strings.ToLower(name)), fmt.Sprintf("<@%s>", slackUserId))
	}

	return text
}
