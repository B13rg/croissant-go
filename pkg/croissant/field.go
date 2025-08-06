// Croissant spec filetypes and relations.
package croissant

type Field struct {
	NType              string   `json:"@type"`
	Source             Source   `json:"source"`
	DataType           DataType `json:"dataType"`
	Repeated           bool     `json:"repeated,omitempty"`
	References         []*Field `json:"references,omitempty"`
	SubField           []*Field `json:"subField,omitempty"`
	ParentField        []*Field `json:"parentField,omitempty"`
	EquivalentProperty string   `json:"equivalentProperty,omitempty"`
}

func NewField() *Field {
	return &Field{}
}

// Type used to group data sources.
type Source struct {
	DataSource *DataSource
	FileObject *FileObject
	FileSet    *FileSet
	RecordSet  *RecordSet
}

func NewSource() *Source {
	return &Source{}
}
