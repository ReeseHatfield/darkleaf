package main

import (
	"fmt"
	"os"

	"github.com/ReeseHatfield/runtime/core"
	"github.com/ReeseHatfield/runtime/fs"
	sec "github.com/ReeseHatfield/runtime/key"
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
	runFlag    bool
	decodeFlag string
)

func main() {

	var key string

	if !isInited() {

		fmt.Println("Projected has NOT already been inited")
		fmt.Println("Setting up darkleaf...")
		key = darkleafInit()
		fmt.Println("Darkleaf setup complete")

	} else {

		serial, err := sec.GetSerialFromUser("Select your key.")
		if err != nil {
			fmt.Errorf("Could not key kehy from user")
			os.Exit(1)
		}

		key = sec.Hash(serial)
	}

	fmt.Printf("key: %v\n", key)

	conf := core.Config{
		Key: key,
	}

	core := core.NewCore(conf)

	core.Run()

}

func darkleafInit() string {
	err := fs.RootMkdirP("$HOME/.darkleaf")

	if err != nil {
		fmt.Println("Could not init darkleaf")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("\033[34m" + "WELCOME TO DARKLEAF")

	serial, err := sec.GetSerialFromUser("Select a new device to use as your key.")
	key := sec.Hash(serial)

	if err != nil {
		fmt.Println("Could not init darkleaf")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return key
}

func isInited() bool {
	exists, err := fs.FileExists("$HOME/.darkleaf")
	if err != nil {
		//probably unsupport os? fixme

		return false
	}

	return exists
}
