package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func killServer(pidFile string) error {

	buff, err := os.ReadFile(pidFile)
	if err != nil {
		return fmt.Errorf("error reading pid file: %v", err)
	}
	if buff == nil || len(buff) < 1 {
		return fmt.Errorf("invalid pid file: %s", pidFile)
	}
	buffString := string(buff[:])
	pid, err := strconv.Atoi(buffString)
	if err != nil {
		return fmt.Errorf("error parsing pid from file: %v", err)
	}
	fmt.Printf("Killing server with pid %v from file %s\n", pid, pidFile)
	return nil
}
func main() {
	if err := killServer("server.pid"); err != nil {
		log.Fatalf("Error: %s\n", err)
	}
}
