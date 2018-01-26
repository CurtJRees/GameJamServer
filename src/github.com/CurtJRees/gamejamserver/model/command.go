package model

import "strconv"

type Command struct {
	X    int    `json:"x"`
	Y    int    `json:"y"`
	Name string `json:"name"`
}

type Commands []Command

func CommandToString(command Command) string {
	return "\tName = " + command.Name + "\tX = " + strconv.Itoa(command.X) + "\tY = " + strconv.Itoa(command.Y)
}