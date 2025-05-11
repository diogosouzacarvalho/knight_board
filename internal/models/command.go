package models

import (
	"strings"
)

type Commands struct {
	Commands []string `json:"commands"`
}

func (c *Commands) IsValid() bool {
	if len(c.Commands) < 1 {
		return false
	}

	return strings.HasPrefix(c.Commands[0], string(CommandTypeStart))
}

type CommandType string

const (
	CommandTypeMove   CommandType = "MOVE"
	CommandTypeRotate CommandType = "ROTATE"
	CommandTypeStart  CommandType = "START"
)

type Direction string

const (
	DirectionNorth Direction = "NORTH"
	DirectionSouth Direction = "SOUTH"
	DirectionEast  Direction = "EAST"
	DirectionWest  Direction = "WEST"
)
