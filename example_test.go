package mqttpattern_test

import "fmt"

import mqttpattern "github.com/amir-yaghoubi/mqtt-pattern"

func ExampleMatches() {
	fmt.Println(mqttpattern.Matches("foo/+/baz", "foo/bar/baz"))
	fmt.Println(mqttpattern.Matches("foo/#", "foo/bar/baz"))
	// Output:
	// true
	// true
}
