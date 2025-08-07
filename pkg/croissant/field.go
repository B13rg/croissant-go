// Croissant spec filetypes and relations.
package croissant

import (
	"github.com/b13rg/croissant-go/pkg/types"
)

type Field struct {
	// Must be field.
	NType string `json:"@type"`
	// Node ID.
	NId string `json:"@id"`
	// Name of the Field.
	Name string `json:"name"`
	// Description of the Field.
	Description string `json:"description"`
	// The data types that correspond to the Field.
	DataType types.StringOrSlice `json:"dataType,omitempty"`
	// The source of data for the field.
	Source DataSource `json:"source,omitempty"`
	// If true the Field is a list of DataType values.
	Repeated bool `json:"repeated,omitempty"`
	// A property URL that is equivalent to this field
	EquivalentProperty string `json:"equivalentProperty,omitempty"`
	// References to other fields in a different RecordSet.
	References ClassRefList `json:"references,omitempty"`
	// List of Fields nested inside this one.
	SubField []Field `json:"subField,omitempty"`
	// References to other Fields in the same RecordSet.
	ParentField ClassRefList `json:"parentField,omitempty"`
}

func NewField() *Field {
	return &Field{}
}

type FieldRef struct {
	Field ClassRefItem `json:"field"`
}
