// Croissant spec filetypes and relations.
package croissant

import "github.com/b13rg/croissant-go/pkg/types"

type RecordSet struct {
	// Must be RecordSet
	NType string `json:"@type"`
	// Node ID
	ClassRefItem
	// Name of the RecordSet
	Name string `json:"name"`
	// Description of the RecordSet
	Description string `json:"description"`
	// The data types that correspond to all fields in the RecordSet.
	DataType types.StringOrSlice `json:"dataType"`
	// One or more Fields that uniquely identify records in the RecordSet.
	Key ClassRefList `json:"key,omitempty"`
	// List of data element Fields that appear in the RecordSet.
	Field []Field `json:"field"`
	// One or more records that constitute the data of the RecordSet.
	Data []interface{} `json:"data,omitempty"`
	// One or more records provided as an example of the RecordSet.
	Examples []interface{} `json:"examples,omitempty"`
}

func NewRecordSet() *RecordSet {
	return &RecordSet{}
}

type DataType struct {
	// MIME type
	DataType string
}

func NewDataType(mimeType string) *DataType {
	return &DataType{DataType: mimeType}
}

type Split struct {
	TrainSplit      string
	TestSplit       string
	ValidationSplit string
}

func NewSplit() *Split {
	return &Split{}
}
