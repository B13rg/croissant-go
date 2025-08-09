// Croissant spec filetypes and relations.
package croissant

import "github.com/b13rg/croissant-go/pkg/types"

type RecordSet struct {
	// Must be RecordSet
	Type string `json:"@type"`
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

func (obj *RecordSet) Validate() ([]types.CroissantWarning, []types.CroissantError) {
	listWarn := []types.CroissantWarning{}
	listError := []types.CroissantError{}

	if obj.Name == "" {
		listWarn = append(listWarn,
			types.CroissantWarning{
				Message: "recordSet Name should be set",
				Value:   obj.ID,
			},
		)
	}
	if obj.Description == "" {
		listWarn = append(listWarn,
			types.CroissantWarning{
				Message: "recordSet description should be set",
				Value:   obj.ID,
			},
		)
	}
	if len(obj.Field) == 0 {
		listWarn = append(listWarn,
			types.CroissantWarning{
				Message: "recordSet specifies no Fields",
				Value:   obj.ID,
			},
		)
	}

	return listWarn, listError
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

func NewSplit(trainSplit string, testSplit string, validationSplit string) *Split {
	return &Split{TrainSplit: trainSplit, TestSplit: testSplit, ValidationSplit: validationSplit}
}
