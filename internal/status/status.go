package status

import "errors"

const StatusSuccess = "SUCCESS"

var (
	ErrInvalidStartPosition = errors.New("INVALID_START_POSITION")
	ErrOutOfBoard           = errors.New("OUT_OF_THE_BOARD")
	ErrGeneric              = errors.New("GENERIC_ERROR")
)
