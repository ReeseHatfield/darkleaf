package main

import (
	"fmt"
	"os"

	"github.com/ReeseHatfield/runtime/fs"
	"github.com/ReeseHatfield/runtime/security"
)

/*

----- DESIGN ------
Flags:
	- init
		setup hw key
		Ask for path for tree fs
			prolly just use /var/darkleaf
	- run (default?)
		Ask for hw key
		needs to remembere where the root is
	- decode [enc-file]
		Ask for hw key



Prolly need a key module?

Will need to spin up sub processes
	- mostly for run

*/

var (
	initFlag   string
	runFlag    string
	decodeFlag string
)

func main() {

	if !security.IsRoot() {
		fmt.Println("deakleaf must be ran as root, try using sudo")
		os.Exit(1)
	}

	if !isInited() {
		fmt.Println("Projected has NOT already been inited")
		fmt.Println("Setting up darkleaf...")
		darkleafInit()
		fmt.Println("Darkleaf setup complete")
	}

}

func darkleafInit() {
	err := fs.RootMkdirP("/var/darkleaf")

	// do any other configuration setup that needs done here
	if err != nil {
		fmt.Println("Could not init darkleaf")
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func isInited() bool {
	exists, err := fs.FileExists("/var/darkleaf")
	if err != nil {
		//probably unsupport os? fixme

		return false
	}

	return exists
}
