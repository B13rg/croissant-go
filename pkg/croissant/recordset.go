// Croissant spec filetypes and relations.
package croissant

type JSON string

type RecordSet struct {
	NType    string   `json:"@type"`
	Field    []Field  `json:"field"`
	Key      []string `json:"key,omitempty"`
	Data     []JSON   `json:"data,omitempty"`
	Examples []JSON   `json:"examples,omitempty"`
	Source   Source   `json:"-"`
}

func NewRecordSet() *RecordSet {
	return &RecordSet{}
}

type DataType struct {
	// MIME type
	DataType string
}

func NewDataType() *DataType {
	return &DataType{}
}

type Split struct {
	TrainSplit      string
	TestSplit       string
	ValidationSplit string
}

func NewSplit() *Split {
	return &Split{}
}
