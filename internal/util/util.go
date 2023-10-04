package util

import (
	"bufio"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func ReadInput() string {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

func StringToMap(data string) map[string]string {
	lines := strings.Split(string(data), "\n")
	configMap := make(map[string]string)

	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			configMap[key] = value
		}
	}

	return configMap
}
