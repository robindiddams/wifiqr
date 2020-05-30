package system

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"

	"github.com/keybase/go-keychain"
)

// GetWifiPassword pulls the wifi password from keychain
func GetWifiPassword(ssid string) (string, error) {
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetAccount(ssid)
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnAttributes(true)
	query.SetReturnData(true)
	results, err := keychain.QueryItem(query)
	if err != nil {
		return "", fmt.Errorf("error getting keychain record: %w", err)
	}
	if len(results) != 1 {
		return "", fmt.Errorf("unable to find password for %s", ssid)
	}
	return string(results[0].Data), nil
}

var airportSSIDRE = regexp.MustCompile("\\s+SSID:\\s([A-z0-9 \\-]+)\\s")

// GetConnectedSSID executes the airport binary to get the connected network
func GetConnectedSSID() (string, error) {
	cmd := exec.Command("/System/Library/PrivateFrameworks/Apple80211.framework/Resources/airport", "-I")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error running command %s: %w", cmd.String(), err)
	}
	matches := airportSSIDRE.FindStringSubmatch(string(output))
	if len(matches) == 2 {
		return matches[1], nil
	}
	return "", errors.New("unable to find ssid")
}

// ViewFile opens the file in preview
func ViewFile(filename string) error {
	cmd := exec.Command("open", filename)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running command %s: %w", cmd.String(), err)
	}
	return nil
}
