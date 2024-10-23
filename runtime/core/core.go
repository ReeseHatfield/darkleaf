package core

import (
	"fmt"
	"log"
	"os/exec"
)

type Core struct {
	key string
}

type Config struct {
	Key string
}

func NewCore(conf Config) *Core {
	c := new(Core)

	c.key = conf.Key

	return c
}

func (c *Core) GetKey() string {
	// enforce some rules about caller?
	return c.key
}

func (c *Core) Run() {
	// async run eventually
	c.run()
}

func (c *Core) run() {

	cmd := exec.Command("/bin/sh", "-c", "cd darkleaf-gui && npm run start --loglevel=verbose")

	cmdOutput, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("Error running command: %v\nOutput: %s", err, string(cmdOutput))
	}

	fmt.Println("Finished?")
	fmt.Println(string(cmdOutput))
}
