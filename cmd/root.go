package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Robindiddams/wifiqr/wifi"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
)

func wifiString(ssid, password string) string {
	return fmt.Sprintf("WIFI:S:%s;T:WPA2;P:%s;;", ssid, password)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wifiqr",
	Short: "create a qr code to log into the wifi you're connected to",
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		if verbose {
			fmt.Println("getting connected network")
		}
		ssid, err := wifi.GetConnectedSSID()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if verbose {
			fmt.Println("connected network is", ssid)
			fmt.Printf("getting %s password, keychain prompt incomming...\n", ssid)
		}
		password, err := wifi.GetWifiPassword(ssid)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fn := fmt.Sprintf("%s.png", strings.ReplaceAll(ssid, " ", "-"))
		if verbose {
			fmt.Printf("got %s password\nwriting qrcode to %s\n", ssid, fn)
		}
		if err := qrcode.WriteFile(wifiString(ssid, password), qrcode.Medium, 256, fn); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if verbose {
			fmt.Printf("qrcode written successfully🎉\n\n\t`open %s`\n", fn)
		} else {
			fmt.Println(fn)
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
	rootCmd.Flags().BoolP("verbose", "v", false, "more words")
}
