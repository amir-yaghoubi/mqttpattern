package mqttpattern_test

import "testing"

import mqttpattern "github.com/amir-yaghoubi/mqtt-pattern"

func TestMattches(t *testing.T) {
	testCases := []struct {
		name    string
		pattern string
		topic   string
		result  bool
	}{
		{name: "Supports patterns with no wildcards", pattern: "foo/bar/baz", topic: "foo/bar/baz", result: true},
		{name: "Doesn't match different pattern and topic", pattern: "foo/bar/baz", topic: "bar/foo/baz", result: false},
		{name: "Supports # at the beginning", pattern: "#", topic: "foo/bar/baz", result: true},
		{name: "Supports # at the end", pattern: "foo/#", topic: "foo/bar/baz", result: true},
		{name: "Supports # at the end and topic has no children", pattern: "foo/bar/#", topic: "foo/bar/baz", result: true},
		{name: "Doesn't support # wildcards with more after them", pattern: "#/foo/bar", topic: "foo/bar/baz", result: false},
		{name: "Supports patterns with + at the beginning", pattern: "+/bar/baz", topic: "foo/bar/baz", result: true},
		{name: "Supports patterns with + at the end", pattern: "foo/bar/+", topic: "foo/bar/baz", result: true},
		{name: "Supports patterns with + at the middle", pattern: "foo/+/baz", topic: "foo/bar/baz", result: true},
		{name: "Supports patterns with multiple wildcards", pattern: "foo/+/#", topic: "foo/bar/baz", result: true},
		{name: "Supports named wildcards", pattern: "foo/+something/#else", topic: "foo/bar/baz", result: true},
		{name: "Supports leading slashes", pattern: "/foo/bar/baz", topic: "/foo/bar/baz", result: true},
		{name: "Supports leading slashes with invalid topic", pattern: "/foo/bar", topic: "/bar/foo", result: false},
	}

	for _, tt := range testCases {
		if mqttpattern.Matches(tt.pattern, tt.topic) != tt.result {
			t.Error(tt.name)
		}
	}
}
