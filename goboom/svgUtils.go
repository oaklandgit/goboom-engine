package goboom

import (
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Point structure
type Pt rl.Vector2

// Command structure
type Command struct {
	Command string
	Params  []string
}

// parsePath function parses the SVG path string into a slice of commands
func parsePath(path string) []Command {
	var commands []Command
	var cmd string
	var params []string
	var param string

	for _, char := range path {
		if strings.ContainsRune("MHVLZL", char) {
			if cmd != "" {
				if param != "" {
					params = append(params, param)
					param = ""
				}
				commands = append(commands, Command{cmd, params})
				params = nil
			}
			cmd = string(char)
		} else if char == ' ' || char == ',' {
			if param != "" {
				params = append(params, param)
				param = ""
			}
		} else {
			param += string(char)
		}
	}
	if cmd != "" {
		if param != "" {
			params = append(params, param)
		}
		commands = append(commands, Command{cmd, params})
	}

	return commands
}

// drawPath function parses and draws an SVG path string
func DrawSVGPath(path string, stroke float32, color rl.Color) (float32, float32) {
	commands := parsePath(path)
	var currentPos, startPos Pt
	var minX, minY, maxX, maxY float32 // need this to calculate the bounding box

	// Initialize bounding box with the first move to command
	if len(commands) > 0 && commands[0].Command == "M" {
		startPos = parseMoveTo(commands[0].Params)
		minX, minY, maxX, maxY = startPos.X, startPos.Y, startPos.X, startPos.Y
		currentPos = startPos
	}

	for _, command := range commands {
		switch command.Command {
		case "M":
			currentPos = parseMoveTo(command.Params)
			startPos = currentPos
		case "H":
			currentPos = parseHorizontalLine(currentPos, command.Params, stroke, color)
		case "V":
			currentPos = parseVerticalLine(currentPos, command.Params, stroke, color)
		case "L":
			currentPos = parseLineTo(currentPos, command.Params, stroke, color)
		case "Z":
			// Closing the path by drawing a line back to the start point
			rl.DrawLineEx(rl.Vector2{X: currentPos.X, Y: currentPos.Y}, rl.Vector2{X: startPos.X, Y: startPos.Y}, stroke, color)
			currentPos = startPos
		}

		// Update the bounding box
		if currentPos.X < minX {
			minX = currentPos.X
		}
		if currentPos.X > maxX {
			maxX = currentPos.X
		}
		if currentPos.Y < minY {
			minY = currentPos.Y
		}
		if currentPos.Y > maxY {
			maxY = currentPos.Y
		}

	}

	// Calculate width and height of the path
	width := maxX - minX
	height := maxY - minY

	return width, height
}

func parseMoveTo(params []string) Pt {
	if len(params) < 2 {
		return Pt{}
	}
	x, _ := strconv.ParseFloat(params[0], 32)
	y, _ := strconv.ParseFloat(params[1], 32)
	return Pt{X: float32(x), Y: float32(y)}
}

func parseHorizontalLine(current Pt, params []string, stroke float32, color rl.Color) Pt {
	if len(params) < 1 {
		return current
	}
	x, _ := strconv.ParseFloat(params[0], 32)
	newPos := Pt{X: float32(x), Y: current.Y}
	// rl.DrawLine(int32(current.X), int32(current.Y), int32(newPos.X), int32(newPos.Y), color)
	rl.DrawLineEx(rl.Vector2{X: current.X, Y: current.Y}, rl.Vector2{X: newPos.X, Y: newPos.Y}, stroke, color)
	return newPos
}

func parseVerticalLine(current Pt, params []string, stroke float32, color rl.Color) Pt {
	if len(params) < 1 {
		return current
	}
	y, _ := strconv.ParseFloat(params[0], 32)
	newPos := Pt{X: current.X, Y: float32(y)}
	// rl.DrawLine(int32(current.X), int32(current.Y), int32(newPos.X), int32(newPos.Y), color)
	rl.DrawLineEx(rl.Vector2{X: current.X, Y: current.Y}, rl.Vector2{X: newPos.X, Y: newPos.Y}, stroke, color)
	return newPos
}

func parseLineTo(current Pt, params []string, stroke float32, color rl.Color) Pt {
	if len(params) < 2 {
		return current
	}
	x, _ := strconv.ParseFloat(params[0], 32)
	y, _ := strconv.ParseFloat(params[1], 32)
	newPos := Pt{X: float32(x), Y: float32(y)}
	// rl.DrawLine(int32(current.X), int32(current.Y), int32(newPos.X), int32(newPos.Y), color)
	rl.DrawLineEx(rl.Vector2{X: current.X, Y: current.Y}, rl.Vector2{X: newPos.X, Y: newPos.Y}, stroke, color)
	return newPos
}
