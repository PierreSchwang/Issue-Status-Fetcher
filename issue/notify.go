package issue

func (c Component) BroadcastUpdate(consumerKey string, consumerSecret string, accessToken string, accessSecret string, webhookUrl string) {
	c.SendToTwitter(consumerKey, consumerSecret, accessToken, accessSecret)
	c.SendToDiscord(webhookUrl)
}
