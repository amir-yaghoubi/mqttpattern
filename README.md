# mqttpattern

Package for matching MQTT patterns with wildcards to extract data from topics

This project is a Go implementation of the [mqtt-pattern](https://github.com/RangerMauve/mqtt-pattern).

## Installation

To install mqttpattern package:

1. Install using go get:

```bash
go get amir-yaghoubi/mqttpattern
```

2. Import it in your code:

```bash
import mqttpattern "github.com/amir-yaghoubi/mqtt-pattern"
```

## Quick start

```go
package main

import (
	"fmt"
	mqttpattern "github.com/amir-yaghoubi/mqtt-pattern"
)

func main() {
	isMatch := mqttpattern.Matches("foo/+/baz", "foo/bar/baz")
	fmt.Printf("isMatch: %t\n", isMatch)
	// isMatch: true

	params := mqttpattern.Extract("foo/+something/baz", "foo/bar/baz")
	fmt.Printf("params: %v\n", params)
	// params: map[something:bar]

	params := mqttpattern.Exec("foo/+something/+otherthing", "foo/bar/baz")
	fmt.Printf("params: %v", params)
	// params: map[otherthing:baz something:bar]

	cleanPattern := mqttpattern.Fill("foo/+bar/#baz", map[string]string{"bar": "BAR", "baz": "BAZ"})
	fmt.Printf("pattern: %s", cleanPattern)
	// pattern: foo/BAR/BAZ

	cleanPattern := mqttpattern.Clean("foo/+something/#otherthing")
	fmt.Printf("pattern: %s", cleanPattern)
	// pattern: foo/+/#
}

```
