package mqttpattern

import (
	"strings"
)

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

// Extract Traverses the pattern and attempts to fetch parameters from the topic.
// Useful if you know in advance that your topic will be valid and want to extract data.
// If the topic doesn't match, or the pattern doesn't contain named wildcards, returns an empty map.
//
// WARNING: Do not use this for validation.
func Extract(pattern string, topic string) map[string]string {
	params := make(map[string]string)
	patternSegments := strings.Split(pattern, seprator)
	topicSegments := strings.Split(topic, seprator)
	topicLen := len(topicSegments)

	for i := range patternSegments {
		if len(patternSegments[i]) == 1 {
			continue
		}

		if topicLen-1 < i {
			break
		}

		pLen := len(patternSegments[i])
		if pLen == 0 {
			continue
		}

		paramName := patternSegments[i][1:pLen]
		if patternSegments[i][0] == all {
			params[paramName] = strings.Join(topicSegments[i:len(topicSegments)], seprator)
			break
		}
		if patternSegments[i][0] == single {
			params[paramName] = topicSegments[i]
		}
	}

	return params
}
