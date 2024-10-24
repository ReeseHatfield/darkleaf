package core

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
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

	keyArg := fmt.Sprintf("--key=%s", c.key)

	cmd := exec.Command("/bin/sh", "-c", "cd darkleaf-gui && npm run start --loglevel=verbose -- "+keyArg)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Error getting stdout pipe: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("Error getting stderr pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatalf("Error starting command: %v", err)
	}

	streamOutput := func(pipe io.ReadCloser, pipeName string) {
		scanner := bufio.NewScanner(pipe)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Printf("[%s] %s\n", pipeName, line)
			if strings.Contains(line, "IPC Message:") {
				fmt.Println("Received message from Electron:", strings.TrimPrefix(line, "IPC Message:"))
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("Error reading from %s: %v", pipeName, err)
		}
	}

	go streamOutput(stdout, "stdout")
	go streamOutput(stderr, "stderr")

	if err := cmd.Wait(); err != nil {
		log.Fatalf("Command finished with error: %v", err)
	}

	fmt.Println("Finished?")
}
