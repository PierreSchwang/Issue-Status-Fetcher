package issue

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
)

func (c Component) SendToTwitter(consumerKey string, consumerSecret string, accessToken string, accessSecret string) {
	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		return
	}
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	var client = twitter.NewClient(httpClient)

	_, _, err := client.Statuses.Update(c.Title+" transitioned to "+c.GetStatus().Name, nil)
	if err != nil {
		log.Fatal("Failed to post twitter status update", err)
	}
}
