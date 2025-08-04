package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	mediaURL  string
	mediaType string
	caption   string
)

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Post an image or video to Instagram",
	Run: func(cmd *cobra.Command, args []string) {
		loadConfig()

		if config.IGID == "" || config.Token == "" || config.Version == "" {
			fmt.Println("❌ Error: config missing. Run `social-cli config set` first.")
			os.Exit(1)
		}
		if mediaURL == "" || caption == "" {
			fmt.Println("❌ Error: --url and --caption are required")
			os.Exit(1)
		}

		creationID, err := createMediaContainer()
		if err != nil {
			log.Fatalf("❌ Failed to create media container: %v", err)
		}
		fmt.Println("✅ Created media container:", creationID)

		// If video, wait for processing before publishing
		if mediaType == "video" {
			fmt.Println("⏳ Waiting for video processing...")
			time.Sleep(20 * time.Second) // crude wait, IG recommends polling
		}

		mediaID, err := publishMedia(creationID)
		if err != nil {
			log.Fatalf("❌ Failed to publish media: %v", err)
		}
		fmt.Println("✅ Published post successfully! Media ID:", mediaID)
	},
}

func init() {
	postCmd.Flags().StringVar(&mediaURL, "url", "", "Media URL (required, must be direct link to image or video file)")
	postCmd.Flags().StringVar(&mediaType, "type", "image", "Media type: image or video (default: image)")
	postCmd.Flags().StringVar(&caption, "caption", "", "Caption text (required)")
	rootCmd.AddCommand(postCmd)
}

func createMediaContainer() (string, error) {
	endpoint := fmt.Sprintf("https://graph.instagram.com/%s/%s/media", config.Version, config.IGID)
	data := url.Values{}

	if mediaType == "video" {
		data.Set("media_type", "REELS")
		data.Set("video_url", mediaURL)
	} else {
		data.Set("image_url", mediaURL)
	}
	data.Set("caption", caption)
	data.Set("access_token", config.Token)

	resp, err := http.PostForm(endpoint, data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("API error: %s", string(body))
	}
	return parseID(body), nil
}

func publishMedia(creationID string) (string, error) {
	endpoint := fmt.Sprintf("https://graph.instagram.com/%s/%s/media_publish", config.Version, config.IGID)
	data := url.Values{}
	data.Set("creation_id", creationID)
	data.Set("access_token", config.Token)

	resp, err := http.PostForm(endpoint, data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("API error: %s", string(body))
	}
	return parseID(body), nil
}

func parseID(body []byte) string {
	var result map[string]string
	if err := json.Unmarshal(body, &result); err == nil {
		if id, ok := result["id"]; ok {
			return id
		}
	}
	return "unknown"
}

