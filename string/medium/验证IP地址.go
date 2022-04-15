package medium

import (
	"strings"
	"regexp"
)

func validIPAddress(queryIP string) string {
	chunkIPv4 := `([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	pattenIPv4 := regexp.MustCompile("^("+chunkIPv4+"\\.){3}" + chunkIPv4 + "$")

	chunkIPv6 := `([0-9a-fA-F]{1,4})`
	pattenIPv6 := regexp.MustCompile("^(" + chunkIPv6 + "\\:){7}" + chunkIPv6 + "$")

	if strings.Contains(queryIP, ".") {
		if pattenIPv4.Match([]byte(queryIP)) {
			return "IPv4"
		} 
	} else if strings.Contains(queryIP, ":") {
		if pattenIPv6.Match([]byte(queryIP)) {
			return "IPv6"
		}
	}
	return "Neither"
}
  