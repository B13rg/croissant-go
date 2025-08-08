package croissant

import (
	"encoding/json"

	"github.com/b13rg/croissant-go/pkg/types"
)

type ClassRefItem struct {
	// ID of the resource.
	ID string `json:"@id,omitempty"`
}

type ClassRefList []ClassRefItem

func (ref *ClassRefList) UnmarshalJSON(data []byte) error {
	// Try unmarshaling as a ClassRefItem
	var single ClassRefItem
	if err := json.Unmarshal(data, &single); err == nil {
		*ref = []ClassRefItem{single}

		return nil
	}
	// Try unmarshaling as a []ClassRefItem
	var multi []ClassRefItem
	if err := json.Unmarshal(data, &multi); err == nil {
		*ref = multi

		return nil
	}
	// Otherwise, error
	return types.CroissantError{
		Message: "StringOrSlice: cannot unmarshal",
		Value:   string(data),
	}
}

func (ref ClassRefList) MarshalJSON() ([]byte, error) {
	switch len(ref) {
	case 0:
		return []byte("{}"), nil
	case 1:
		return json.Marshal(ref[0])
	default:
		return json.Marshal(ref)
	}
}
