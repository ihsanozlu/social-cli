package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type Config struct {
	IGID    string `json:"ig_id"`
	Token   string `json:"token"`
	Version string `json:"version"`
}

var config Config

var (
	igID    string
	token   string
	version string
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage Instagram credentials",
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set Instagram credentials (all or one by one)",
	Run: func(cmd *cobra.Command, args []string) {
		loadConfig() // load existing config if any

		// Update only provided flags (allows one-by-one setting)
		if igID != "" {
			config.IGID = igID
		}
		if token != "" {
			config.Token = token
		}
		if version != "" {
			config.Version = version
		}

		if config.IGID == "" || config.Token == "" || config.Version == "" {
			fmt.Println("⚠️ Warning: config saved but missing required fields (ig-id, token, version)")
		}

		saveConfig(config)
		fmt.Println("✅ Config updated!")
	},
}

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Show current config",
	Run: func(cmd *cobra.Command, args []string) {
		loadConfig()
		fmt.Println("📌 Current config:")
		fmt.Printf(" IG ID   : %s\n", config.IGID)
		fmt.Printf(" Token   : %s\n", maskToken(config.Token))
		fmt.Printf(" Version : %s\n", config.Version)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete config file",
	Run: func(cmd *cobra.Command, args []string) {
		path := getConfigPath()
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("❌ No config file found to delete")
			return
		}
		err := os.Remove(path)
		if err != nil {
			fmt.Println("❌ Failed to delete config:", err)
			return
		}
		fmt.Println("🗑️ Config deleted successfully")
	},
}

func init() {
	// Flags for setting config
	setCmd.Flags().StringVar(&igID, "ig-id", "", "Instagram User ID")
	setCmd.Flags().StringVar(&token, "token", "", "Access Token")
	setCmd.Flags().StringVar(&version, "version", "", "API version (default v23.0)")

	configCmd.AddCommand(setCmd)
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(configCmd)
}

func saveConfig(cfg Config) {
	path := getConfigPath()
	os.MkdirAll(filepath.Dir(path), 0755)
	data, _ := json.MarshalIndent(cfg, "", "  ")
	_ = ioutil.WriteFile(path, data, 0644)
}

func loadConfig() {
	path := getConfigPath()
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return // no config yet
	}
	_ = json.Unmarshal(data, &config)
}

func getConfigPath() string {
	dir, _ := os.UserHomeDir()
	return filepath.Join(dir, ".social-cli", "config.json")
}

// Mask token for display
func maskToken(token string) string {
	if len(token) < 8 {
		return token
	}
	return token[:4] + "..." + token[len(token)-4:]
}

