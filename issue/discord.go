package issue

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
)

func (c Component) SendToDiscord(webhookUrl string) {
	if webhookUrl == "" {
		return
	}
	color, err := strconv.ParseInt(c.GetStatus().Color, 16, 64)
	if err != nil {
		log.Fatal("Failed to convert hex color to int color", err)
		return
	}
	_, err = http.Post(webhookUrl, "application/json", bytes.NewBufferString(
		`{
			  "content": null,
			  "embeds": [
				{
				  "title": "Status Update",
				  "color": `+strconv.FormatInt(color, 10)+`,
				  "fields": [
					{
					  "name": "`+c.Title+`",
					  "value": "`+c.Title+` transitioned to `+c.GetStatus().Name+`"
					}
				  ]
				}
			  ]
			}`))
	if err != nil {
		log.Fatal("Failed to post status update to discord webhook", err)
		return
	}
}
