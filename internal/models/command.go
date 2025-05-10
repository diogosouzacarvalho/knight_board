package models

type Commands struct {
	Commands []string `json:"commands"`
}

type Command string

const (
	CommandMove   Command = "MOVE"
	CommandRotate Command = "ROTATE"
	CommandStart  Command = "START"
)

type Direction string

const (
	DirectionNorth Direction = "NORTH"
	DirectionSouth Direction = "SOUTH"
	DirectionEast  Direction = "EAST"
	DirectionWest  Direction = "WEST"
)
