package util

import "github.com/b13rg/croissant-go/pkg/types"

// If err != nil, wraps it in a Kr8Error with the message.
func ErrorIfCheck(message string, err error) *types.CroissantError {
	if err != nil {
		return &types.CroissantError{Message: message, Value: err}
	}

	return nil
}
