package backend

import (
	"fmt"
	"net"
)

func Telnet(serverIP string) (string, bool) {

	_, err := net.Dial("tcp", serverIP)

	if nil != err {
		telnetResponds := fmt.Sprintf("Something went wrong to connect: %s Error: %s", serverIP, err)

		return telnetResponds, false
	} else {
		telnetResponds := fmt.Sprintf("successfully connected to %s ", serverIP)

		return telnetResponds, true
	}

}
