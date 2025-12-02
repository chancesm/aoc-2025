package util

import "strings"

func NonEmptyLines(s string) []string {
	raw := strings.Split(strings.TrimSpace(s), "\n")
	out := make([]string, 0, len(raw))
	for _, line := range raw {
		line = strings.TrimSpace(line)
		if line != "" {
			out = append(out, line)
		}
	}
	return out
}
func CommaSeparated(s string) []string {
	return strings.Split(s, ",")
}