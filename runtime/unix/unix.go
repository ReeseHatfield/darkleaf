package security

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

// isRoot return whether the uid is root
// https://stackoverflow.com/questions/29733575/how-to-find-the-user-that-executed-a-program-as-root-using-golang
func IsRoot() bool {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("[isRoot] Unable to get current user: %s", err)
	}

	return currentUser.Username == "root"
}

func Lsusb() []string {
	stdout := RunCommand("lsusb")

	lines := strings.Split(stdout, "\n")

	return lines
}

func GetSerial(bus string, dev string) string {
	devicePath := fmt.Sprintf("/dev/bus/usb/%s/%s", bus, dev)

	udevOut := UdevadmInfo(devicePath)

	grepOut := Grep(udevOut, "ID_SERIAL=")

	stdout := RunCommand(fmt.Sprintf("echo \"%s\" | awk -F'=' '{print $2}'", grepOut))

	return stdout

}

func UdevadmInfo(name string) string {

	stdout := RunCommand(fmt.Sprintf("udevadm info --name=%s", name))

	return stdout
}

// serial=$(udevadm info --name=/dev/bus/usb/$bus/$device | grep 'ID_SERIAL=' | awk -F'=' '{print $2}')

func Grep(pipe string, pattern string) string {

	stdout := RunCommand(fmt.Sprintf("echo \"%s\" | grep %s", pipe, pattern))

	return stdout
}

// RunCommand runs a shell command and returns its output as a string
func RunCommand(cmdString string) string {
	cmd := exec.Command("sh", "-c", cmdString)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error executing command '%s': %v\n", cmdString, err)
		os.Exit(1)
	}
	return string(stdout)
}
