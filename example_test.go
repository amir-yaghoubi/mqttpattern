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

func ExampleExtract() {
	params := mqttpattern.Extract("foo/+something/+otherthing", "foo/bar/baz")
	fmt.Printf("%v", params)

	// Output:
	// map[otherthing:baz something:bar]
}

func ExampleExec() {
	params := mqttpattern.Exec("foo/+something/+otherthing", "foo/bar/baz")
	fmt.Printf("%v", params)

	// Output:
	// map[otherthing:baz something:bar]
}

func ExampleFill() {
	cleanPattern := mqttpattern.Fill("foo/+bar/#baz", map[string]string{"bar": "BAR", "baz": "BAZ"})
	fmt.Printf("%v", cleanPattern)

	// Output:
	// foo/BAR/BAZ
}

func ExampleClean() {
	cleanPattern := mqttpattern.Clean("foo/+something/#otherthing")
	fmt.Printf("%v", cleanPattern)

	// Output:
	// foo/+/#
}
