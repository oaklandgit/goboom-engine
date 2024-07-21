package goboom

import (
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Point structure
type Point2 struct {
	X, Y float32
}

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
func DrawSVGPath(path string, stroke float32, color rl.Color) {
	commands := parsePath(path)
	var currentPos, startPos Point2

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

	}
}

func parseMoveTo(params []string) Point2 {
	if len(params) < 2 {
		return Point2{}
	}
	x, _ := strconv.ParseFloat(params[0], 32)
	y, _ := strconv.ParseFloat(params[1], 32)
	return Point2{X: float32(x), Y: float32(y)}
}

func parseHorizontalLine(current Point2, params []string, stroke float32, color rl.Color) Point2 {
	if len(params) < 1 {
		return current
	}
	x, _ := strconv.ParseFloat(params[0], 32)
	newPos := Point2{X: float32(x), Y: current.Y}
	// rl.DrawLine(int32(current.X), int32(current.Y), int32(newPos.X), int32(newPos.Y), color)
	rl.DrawLineEx(rl.Vector2{X: current.X, Y: current.Y}, rl.Vector2{X: newPos.X, Y: newPos.Y}, stroke, color)
	return newPos
}

func parseVerticalLine(current Point2, params []string, stroke float32, color rl.Color) Point2 {
	if len(params) < 1 {
		return current
	}
	y, _ := strconv.ParseFloat(params[0], 32)
	newPos := Point2{X: current.X, Y: float32(y)}
	// rl.DrawLine(int32(current.X), int32(current.Y), int32(newPos.X), int32(newPos.Y), color)
	rl.DrawLineEx(rl.Vector2{X: current.X, Y: current.Y}, rl.Vector2{X: newPos.X, Y: newPos.Y}, stroke, color)
	return newPos
}

func parseLineTo(current Point2, params []string, stroke float32, color rl.Color) Point2 {
	if len(params) < 2 {
		return current
	}
	x, _ := strconv.ParseFloat(params[0], 32)
	y, _ := strconv.ParseFloat(params[1], 32)
	newPos := Point2{X: float32(x), Y: float32(y)}
	// rl.DrawLine(int32(current.X), int32(current.Y), int32(newPos.X), int32(newPos.Y), color)
	rl.DrawLineEx(rl.Vector2{X: current.X, Y: current.Y}, rl.Vector2{X: newPos.X, Y: newPos.Y}, stroke, color)
	return newPos
}
