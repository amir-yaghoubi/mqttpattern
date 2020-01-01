package mqttpattern

import (
	"strings"
)

const seprator = "/"
const single = '+'
const all = '#'

// Exec Validates that topic fits the pattern and parses out any parameters. If the topic doesn't match, it returns empty map
func Exec(pattern string, topic string) map[string]string {
	if Matches(pattern, topic) {
		return Extract(pattern, topic)
	}
	return nil
}

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

func Fill(pattern string, params map[string]string) string {
	patternSegments := strings.Split(pattern, seprator)
	segments := make([]string, 0, len(patternSegments))

	for i := range patternSegments {
		patternParam := patternSegments[i][1:len(patternSegments[i])]
		paramValue, ok := params[patternParam]

		if patternSegments[i][0] == all {
			if ok {
				segments = append(segments, paramValue)
			} else {
				segments = append(segments, "")
			}
			break
		} else if patternSegments[i][0] == single {
			if ok {
				segments = append(segments, paramValue)
			} else {
				segments = append(segments, "")
			}
		} else {
			segments = append(segments, patternSegments[i])
		}
	}

	return strings.Join(segments, seprator)
}

// Clean Removes the named parameters from a pattern.
func Clean(pattern string) string {
	patternSegments := strings.Split(pattern, seprator)
	cleanedSegments := make([]string, 0, len(patternSegments))

	for i := range patternSegments {
		if patternSegments[i][0] == all {
			cleanedSegments = append(cleanedSegments, string(all))
		} else if patternSegments[i][0] == single {
			cleanedSegments = append(cleanedSegments, string(single))
		} else {
			cleanedSegments = append(cleanedSegments, patternSegments[i])
		}
	}

	return strings.Join(cleanedSegments, seprator)
}
