package utils

import (
	"encoding/json"
	"fmt"
)

// PrettyPrint prints a struct into a human readable JSON format
func PrettyPrint(s interface{}) {
	b, _ := json.MarshalIndent(s, "", "  ")

	fmt.Println(string(b))
}
