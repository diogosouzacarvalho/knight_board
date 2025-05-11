package utils

import (
	"encoding/json"
	"fmt"
	"io"
)

func OutputJSON(writer io.Writer, data any) {
	bytes, _ := json.MarshalIndent(data, "", "\t")
	fmt.Fprintln(writer, string(bytes))
}
