# mqttpattern

Package for matching MQTT patterns with wildcards to extract data from topics

This project is a go implementation of the [mqtt-pattern](https://github.com/RangerMauve/mqtt-pattern).

## Example:

```go
package main

import mqttpattern "github.com/amir-yaghoubi/mqtt-pattern"

import "fmt"

func main() {
	isMatch := mqttpattern.Matches("foo/+/baz", "foo/bar/baz")
	fmt.Printf("isMatch: %t\n", isMatch)
	// isMatch: true

	params := mqttpattern.Extract("foo/+something/baz", "foo/bar/baz")
	fmt.Printf("params: %v\n", params)
	// params: map[something:bar]
}

```
