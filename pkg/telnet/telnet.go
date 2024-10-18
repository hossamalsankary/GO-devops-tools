package telnet

import (
	"fmt"
	"net"
	"time"
)

func Telnet(serverIP string) (string, bool) {
	timeout := 5 * time.Second
	conn, err := net.DialTimeout("tcp", serverIP, timeout)

	// If there's an error during connection
	if err != nil {
		return fmt.Sprintf("Something went wrong while connecting to %s. Error: %v", serverIP, err), false
	}

	defer conn.Close()

	// If connection is successful
	return fmt.Sprintf("Successfully connected to %s", serverIP), true
}
