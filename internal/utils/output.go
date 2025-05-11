package utils

import (
	"encoding/json"
	"fmt"
)

func OutputJSON(data any) {
	bytes, _ := json.MarshalIndent(data, "", "\t")
	fmt.Println(string(bytes))
}
