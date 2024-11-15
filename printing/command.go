package printing

import (
	"encoding/hex"
	"regexp"
	"strings"
)

func HandleCommand(layout string) ([]byte, string, error) {
	var command []byte
	reCommand := regexp.MustCompile(`(?s)<EC>(.*?)<DC>`)
	matches := reCommand.FindStringSubmatch(layout)
	if len(matches) > 1 {
		commandList := strings.Split(matches[1], "\n")
		for _, c := range commandList {
			c = strings.TrimPrefix(c, "<XCT>")
			byteList := strings.Split(c, ",")
			for _, bStr := range byteList {
				bStr = strings.TrimSpace(bStr)
				if len(bStr)%2 != 0 {
					bStr = "0" + bStr
				}
				b, err := convertHexToBytes(bStr)
				if err != nil {
					return nil, layout, err
				}
				command = append(command, b...)
			}
		}
		layout = strings.ReplaceAll(layout, matches[0], "")
	}
	return command, layout, nil
}

func convertHexToBytes(hexStr string) ([]byte, error) {
	data, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	return data, nil
}
