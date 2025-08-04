package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get user posts or post info",
	Run: func(cmd *cobra.Command, args []string) {
		loadConfig()

		endpoint := fmt.Sprintf("https://graph.instagram.com/%s/%s/media?fields=id,caption,media_url,timestamp&access_token=%s",
			config.Version, config.IGID, config.Token)

		resp, err := http.Get(endpoint)
		if err != nil {
			fmt.Println("❌ Error fetching posts:", err)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

