package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "social-cli",
	Short: "SocialCLI is a tool to interact with Instagram Graph API",
	Long:  "SocialCLI lets you configure your Instagram credentials, publish posts, and fetch info using Instagram Graph API.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

