package croissant

import (
	"encoding/json"

	"github.com/b13rg/croissant-go/pkg/types"
)

type ClassRefItem struct {
	Id string `json:"@id"`
}

type ClassRefList struct {
	Items []ClassRefItem
}

func (ref *ClassRefList) UnmarshalJSON(data []byte) error {
	// Try unmarshaling as a ClassRefItem
	var single ClassRefItem
	if err := json.Unmarshal(data, &single); err == nil {
		ref.Items = []ClassRefItem{single}

		return nil
	}
	// Try unmarshaling as a []ClassRefItem
	var multi []ClassRefItem
	if err := json.Unmarshal(data, &multi); err == nil {
		ref.Items = multi

		return nil
	}
	// Otherwise, error
	return types.CroissantError{
		Message: "StringOrSlice: cannot unmarshal",
		Value:   string(data),
	}
}
