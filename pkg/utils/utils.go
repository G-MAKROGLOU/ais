package utils

import (
	"encoding/json"
	"fmt"
)

// PrettyPrint prints a struct into a pretty JSON format
func PrettyPrint(s interface{}) {
	b, _ := json.MarshalIndent(s, "", "  ")

	fmt.Println(string(b))
}
