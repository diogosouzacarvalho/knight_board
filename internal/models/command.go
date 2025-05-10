package models

type Commands struct {
	Commands []string `json:"commands"`
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
