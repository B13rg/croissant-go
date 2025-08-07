// Place internal types here.
package types

import "fmt"

type CroissantError struct {
	// Message to show the user.
	Message string
	// Value to include with message
	Value any
}

func (e CroissantError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Value)
}
