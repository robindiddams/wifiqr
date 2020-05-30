package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Robindiddams/wifiqr/system"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
)

func wifiString(ssid, password string) string {
	return fmt.Sprintf("WIFI:S:%s;T:WPA2;P:%s;;", ssid, password)
}

func qrFilename(ssid string) string {
	return fmt.Sprintf("%s.png", strings.ReplaceAll(ssid, " ", "-"))
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wifiqr",
	Short: "create a qr code to log into the wifi you're connected to",
	Run: func(cmd *cobra.Command, args []string) {
		ssid, err := system.GetConnectedSSID()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		password, err := system.GetWifiPassword(ssid)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fn := filepath.Join(os.TempDir(), qrFilename(ssid))
		if err := qrcode.WriteFile(wifiString(ssid, password), qrcode.Medium, 256, fn); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := system.ViewFile(fn); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// rootCmd.Flags().BoolP("verbose", "v", false, "more words") // TODO
}
