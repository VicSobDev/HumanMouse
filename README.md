# HumanMouse

A Go package that simulates random human-like mouse movements, allowing for custom movement bounds.

## Installation

```bash
go get -u github.com/VicSobDev/HumanMouse
```

## Features
- Simulate human-like mouse movements with user-defined bounds.
- Generate sequences of mouse movements to mimic real-world user behavior.
- Easy to integrate and customizable.


## Sample Outputs

**Brave's Mouse Movements:** 

```
x: 271, y: 766, duration: 14149ms
x: 271.90234375, y: 766.05859375, duration: 1
x: 272.10546875, y: 765.8515625, duration: 2
x: 272.91796875, y: 765.64453125, duration: 6
x: 272.91796875, y: 765.4375, duration: 6
x: 272.91796875, y: 765.4375, duration: 7
x: 273.12109375, y: 765.4375, duration: 7
x: 273.12109375, y: 765.4375, duration: 8
x: 273.52734375, y: 765.4375, duration: 10
x: 273.52734375, y: 765.4375, duration: 11
x: 273.73046875, y: 765.4375, duration: 11
x: 273.73046875, y: 765.4375, duration: 12
x: 273.73046875, y: 765.23046875, duration: 16
x: 274.74609375, y: 765.0234375, duration: 17
(etc..)
```

**HumanMouse's Generated Movements:**

```
X: 271, Y: 764.4375, Duration: 1
X: 275.6875, Y: 764.4375, Duration: 2
X: 275.6875, Y: 760.921875, Duration: 4
X: 275.6875, Y: 760.921875, Duration: 5
X: 278.421875, Y: 760.921875, Duration: 6
X: 278.421875, Y: 756.234375, Duration: 7
X: 278.8125, Y: 756.234375, Duration: 7
X: 281.546875, Y: 756.234375, Duration: 8
X: 281.546875, Y: 753.890625, Duration: 13
X: 286.625, Y: 753.890625, Duration: 14
X: 286.625, Y: 751.546875, Duration: 15
X: 290.921875, Y: 751.546875, Duration: 16
X: 290.921875, Y: 748.421875, Duration: 17
X: 294.046875, Y: 748.421875, Duration: 18
(etc..)
```

## Usage

Define your bounds, decide if the start is random, and generate mouse movements:

```go
package main

import (
	"fmt"
	"github.com/VicSobDev/HumanMouse"
)

func main() {
	m := &HumanMouse.Trajectory{
		IsStartRandom: true,
		Bounds: HumanMouse.Bounds{
			MinX: 100.0,
			MaxX: 800.0,
			MinY: 50.0,
			MaxY: 600.0,
		},
	}
	movements := m.GenerateMouseMovements()
	for _, move := range movements {
		fmt.Println(move)
	}
}
```
