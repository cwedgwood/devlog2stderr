//

package main

import (
	"bytes"
	"log"
	"strconv"
	"strings"
)

// "<15>Aug 16 04:35:01 root[4567]: p7"
// "<p>mmm dd hh:mm:ss MESSAGE"
func rfc3164(b []byte) bool {
	if len(b) < 18 || b[0] != '<' {
		return false
	}
	var e int
	for e = 1; e <= 5; e++ {
		if b[e] == '>' {
			break
		}
	}
	if b[e] != '>' {
		return false
	}
	pf, err := strconv.Atoi(string(b[1:e]))
	if err != nil {
		return false
	}
	// discard <..>
	b = b[e+1:]

	// skip over date, count 3 spaces, should happen inside of 15 bytes, allow for 25, assume we don't get 2+ adjacent spaces
	// because if we did other things would break already
	skipSpace := 3
	for e = 0; e < 25; e++ {
		if e >= len(b) {
			return false
		}
		if b[e] == ' ' {
			skipSpace--
			if skipSpace == 0 {
				break
			}
		}
	}
	if skipSpace != 0 {
		return false
	}
	// discard date
	b = b[e+1:]

	// trim and print
	b = bytes.TrimRight(b, "\r\n\t ")
	var f string
	if pf>>3 < len(fstr) {
		f = fstr[pf>>3]
	}
	if f == "" {
		f = strconv.Itoa(pf >> 3)
	}
	log.Printf("[%-6s] [%-5s] \t%s", f, pstr[pf&7], strconv.Quote(string(b)))
	return true
}

// see syslog.h, made some shorter
var pstr = strings.Fields("emerg alert crit error warn ntce info debug")
var fstr = strings.Fields("kern user mail daemon auth syslog lpr news uucp cron athprv ftp 12 13 14 15 local0 local1 local2 local3 local4 local5 local6 local7")
