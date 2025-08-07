// Place internal types here.
package types

import (
	"encoding/json"
	"fmt"
)

type StringOrSlice []string

func (s *StringOrSlice) UnmarshalJSON(data []byte) error {
	// Try unmarshaling as a string
	var single string
	if err := json.Unmarshal(data, &single); err == nil {
		*s = []string{single}

		return nil
	}
	// Try unmarshaling as a []string
	var multi []string
	if err := json.Unmarshal(data, &multi); err == nil {
		*s = multi

		return nil
	}
	// Otherwise, error
	return CroissantError{
		Message: "StringOrSlice: cannot unmarshal",
		Value:   string(data),
	}
}

type CroissantError struct {
	// Message to show the user.
	Message string
	// Value to include with message
	Value any
}

func (e CroissantError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Value)
}
