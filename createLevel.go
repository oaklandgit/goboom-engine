package main

import (
	"fmt"
	"strings"
)

func CreateLevel(
	name string,
	// size rl.Vector2,
	levelMap string,
	cellW, cellH float32,
	levelLookup map[rune]func() *GameObj) *GameObj {

	level := NewGameObject(name)
	levelRows := strings.Split(levelMap, "\n")

	fmt.Println("level", level)
	
	for row := range levelRows {
		for col, char := range levelRows[row] {
			if objFunc, ok := levelLookup[char]; ok {
				obj := objFunc()
				obj.Position.X = float32(col) * cellW
				obj.Position.Y = float32(row) * cellH
				// obj.Parent = level
				// level.Children = append(level.Children, obj)
				level.AddChildren(obj)
				// level.Size = size
			}
		}
		
	}

	return level
}