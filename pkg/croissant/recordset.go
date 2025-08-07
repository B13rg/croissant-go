// Croissant spec filetypes and relations.
package croissant

type JSON string

type RecordSet struct {
	// Must be RecordSet
	NType string `json:"@type"`
	// Node ID
	NId string `json:"@id"`
	// List of data element Fields that appear in the RecordSet.
	Field []Field `json:"field"`
	// One or more Fields that uniquely identify records in the RecordSet.
	Key []string `json:"key,omitempty"`
	// One or more records that constitute the data of the RecordSet.
	Data []JSON `json:"data,omitempty"`
	// One or more records provided as an example of the RecordSet.
	Examples []JSON `json:"examples,omitempty"`
	//
	Source Source `json:"-"`
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
