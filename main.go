// listen for syslog messages, send to stderr (log.Print)

package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const sockpath = "/dev/log"

func dgram() {
	dgsock, err := net.ListenPacket("unixgram", sockpath)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	pktbuf := make([]byte, 8192)
	// loop getting packets
	for {
		n, _, err := dgsock.ReadFrom(pktbuf)
		if err != nil {
			log.Printf("ERROR reading packet: '%v'", err)
			continue // FIXME, continue on error isn't always right
		}

		// rfc3164?
		if rfc3164(pktbuf[:n]) {
			continue
		}

		// best effort
		log.Println(strconv.Quote(strings.TrimRight(string(pktbuf[:n]), "\r\n\t ")))
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	if err := os.RemoveAll(sockpath); err != nil {
		log.Fatal(err)
	}
	dgram()
}
