package srttotext

import (
	"bufio"
	"strings"
	"unicode"
)

func Convert(srtContent string) string {
	scanner := bufio.NewScanner(strings.NewReader(srtContent))

	var result strings.Builder

	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line contains timing information
		if strings.Contains(line, "-->") {
			continue
		}

		if onlyDigitsOrWhitespace(line) {
			continue
		}

		if len(strings.TrimSpace(line)) > 0 {
			result.WriteString(line + " ")
		}
	}

	return strings.TrimSpace(result.String())
}

func onlyDigitsOrWhitespace(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) && !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}
