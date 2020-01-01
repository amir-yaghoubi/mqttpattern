package mqttpattern

import "strings"

const seprator = "/"

const single = '+'
const all = '#'

// Matches validates whether topic fits the pattern or not
func Matches(pattern string, topic string) bool {
	patternSegments := strings.Split(pattern, seprator)
	topicSegments := strings.Split(topic, seprator)

	patternLen := len(patternSegments)
	topicLen := len(topicSegments)
	lastIndex := patternLen - 1

	for i := range patternSegments {
		tLen := len(topicSegments[i])

		if len(patternSegments[i]) == 0 && tLen == 0 {
			continue
		}
		if tLen == 0 && patternSegments[i][0] != all {
			return false
		}
		if patternSegments[i][0] == all {
			return i == lastIndex
		}
		if patternSegments[i][0] != single && patternSegments[i] != topicSegments[i] {
			return false
		}
	}

	return patternLen == topicLen
}
