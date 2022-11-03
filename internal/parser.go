package pwrstat_exporter

import (
	"fmt"
	"strings"
)

func parse(input string) map[string]string {
	measures := make(map[string]string)

	splits := strings.Split(input, "\n")

	// remove header element
	skipFirst := strings.EqualFold(splits[0], "STATUS")
	if skipFirst {
		splits = splits[1:]
	}

	// clean line and split into key/val
	for _, line := range splits {
		cleanLine := strings.TrimSpace(line)
		lineSplit := strings.Split(cleanLine, "=")
		measures[lineSplit[0]] = lineSplit[1]
	}

	for key, val := range measures {
		fmt.Println(key, `:`, val)
	}

	return measures
}
