package security

import (
	"log"
	"os/user"
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
