// Croissant spec filetypes and relations.
package croissant

type Field struct {
	// Must be field
	NType string `json:"@type"`
	// Node ID
	NId string `json:"@id"`
	// The source of data for the field.
	Source Source `json:"source"`
	// The data types that correspond to the Field.
	DataType []*DataType `json:"dataType"`
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

// Type used to group data sources.
type Source struct {
	DataSource ClassRefItem
	FileObject ClassRefItem
	FileSet    ClassRefItem
	RecordSet  ClassRefItem
}

func NewSource() *Source {
	return &Source{}
}
