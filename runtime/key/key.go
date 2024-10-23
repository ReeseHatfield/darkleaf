package key

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"

	nix "github.com/ReeseHatfield/runtime/unix"
	"github.com/manifoldco/promptui"
)

func GetSerialFromUser(message string) (string, error) {
	options := nix.Lsusb()

	prompt := promptui.Select{
		Label: "\033[34m" + message,
		Items: options,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", fmt.Errorf("Could not prompt user for key")
	}

	bus, dev, err := parseBusAndDevice(result)

	if err != nil {
		return "", fmt.Errorf("Could not parse device")
	}

	serial := nix.GetSerial(bus, dev)

	return serial, nil
}

func parseBusAndDevice(input string) (string, string, error) {
	re := regexp.MustCompile(`Bus (\d+) Device (\d+):`)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 3 {
		return "", "", fmt.Errorf("invalid input format")
	}

	return matches[1], matches[2], nil
}

// may wanna swap this to bytes depending on usage
func Hash(input string) string {
	hash := sha256.New()

	hash.Write([]byte(input))

	hashBytes := hash.Sum(nil)

	hashString := hex.EncodeToString(hashBytes)

	return hashString
}
