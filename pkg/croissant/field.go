// Croissant spec filetypes and relations.
package croissant

import (
	"encoding/json"

	"github.com/b13rg/croissant-go/pkg/types"
)

type Field struct {
	// Must be field.
	NType string `json:"@type"`
	// Node ID.
	ClassRefItem
	// Name of the Field.
	Name string `json:"name"`
	// Description of the Field.
	Description string `json:"description"`
	// The data types that correspond to the Field.
	DataType types.StringOrSlice `json:"dataType,omitempty"`
	// The source of data for the field.
	Source *DataSource `json:"source,omitempty"`
	// If true the Field is a list of DataType values.
	Repeated bool `json:"repeated,omitempty"`
	// A property URL that is equivalent to this field
	EquivalentProperty string `json:"equivalentProperty,omitempty"`
	// References to other fields in a different RecordSet.
	References FieldRefSlice `json:"references,omitempty"`
	// List of Fields nested inside this one.
	SubField []Field `json:"subField,omitempty"`
	// References to other Fields in the same RecordSet.
	ParentField FieldRefSlice `json:"parentField,omitempty"`
}

func NewField() *Field {
	return &Field{}
}

type FieldRef struct {
	Field ClassRefItem `json:"field"`
}

type FieldRefSlice []FieldRef

func (ref *FieldRefSlice) UnmarshalJSON(data []byte) error {
	// Try unmarshaling as a FieldRef
	var single FieldRef
	if err := json.Unmarshal(data, &single); err == nil {
		*ref = []FieldRef{single}

		return nil
	}
	// Try unmarshaling as a []FieldRef
	var multi []FieldRef
	if err := json.Unmarshal(data, &multi); err == nil {
		*ref = multi

		return nil
	}
	// Otherwise, error
	return types.CroissantError{
		Message: "FieldRefSlice: cannot unmarshal",
		Value:   string(data),
	}
}

func (ref FieldRefSlice) MarshalJSON() ([]byte, error) {
	switch len(ref) {
	case 0:
		return []byte("{}"), nil
	case 1:
		return json.Marshal(ref[0])
	default:
		return json.Marshal(ref)
	}
}
