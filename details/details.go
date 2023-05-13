package details

import (
	"log"
	"net"
	"os"
)

func GetHostName() (string, error) {
	//we're gonna use OS package, we'll have something called host name in package OS
	return os.Hostname()
}

func GetIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
