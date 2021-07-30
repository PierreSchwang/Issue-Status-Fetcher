package cmd

import (
	"github.com/spf13/cobra"
	"issue-status-fetcher/issue"
	"log"
	"time"
)

type Options struct {
	Url                       string
	DiscordWebhookUrl         string
	TwitterConsumerKey        string
	TwitterConsumerSecret     string
	TwitterAccessToken        string
	TwitterAccessSecret       string
	GitHubPersonalAccessToken string
}

var options = &Options{}

func Run() {
	cmd := cobra.Command{}
	cmd.Flags().StringVar(&options.Url, "url", "PierreSchwang/status.gommehd.net", "status page repository")
	cmd.Flags().StringVar(&options.DiscordWebhookUrl, "discord", "", "discord webhook url")
	cmd.Flags().StringVar(&options.TwitterConsumerKey, "twitterConsumerKey", "", "twitter app api key")
	cmd.Flags().StringVar(&options.TwitterConsumerSecret, "twitterConsumerSecret", "", "twitter app api secret")
	cmd.Flags().StringVar(&options.TwitterAccessToken, "twitterAccessToken", "", "twitter app access token")
	cmd.Flags().StringVar(&options.TwitterAccessSecret, "twitterAccessSecret", "", "twitter app access secret")
	cmd.Flags().StringVar(&options.GitHubPersonalAccessToken, "githubPersonalToken", "", "GitHub Personal Access Token")
	err := cmd.Execute()
	if err != nil {
		log.Fatal("Failed to execute command", err)
		return
	}

	// GH Ratelimit:
	// Authenticated: 	5000 / hour
	// Unauthenticated: 60 / Hour
	duration := 2 * time.Minute
	if options.GitHubPersonalAccessToken != "" {
		duration = 10 * time.Second
	}
	Update(duration)
}

func Update(untilNext time.Duration) {
	for true {
		log.Println("Reloading status page data")
		components, err := issue.Components(options.Url, options.GitHubPersonalAccessToken)
		if err != nil {
			log.Fatal("Failed to get all components ", err)
			return
		}
		components = issue.RefreshComponents(components)

		for _, component := range components {
			component.BroadcastUpdate(
				options.TwitterConsumerKey, options.TwitterConsumerSecret, options.TwitterAccessToken, options.TwitterAccessSecret,
				options.DiscordWebhookUrl,
			)
		}
		time.Sleep(untilNext)
	}
}
