// Croissant spec filetypes and relations.
package croissant

type DataSource struct {
	NType      string      `json:"@type"`
	FileObject *FileObject `json:"fileObject"`
	FileSet    *FileSet    `json:"fileSet"`
	RecordSet  *RecordSet  `json:"recordSet"`
	Extract    Extract     `json:"extract"`
	Transform  Transform   `json:"transform"`
	Format     Format      `json:"format"`
}

func NewDataSource() *DataSource {
	return &DataSource{}
}

type Format struct{}

type Extract struct {
	FileProperty ContentExtractionEnumeration
	Column       string `json:"column"`
	JsonPath     string `json:"jsonPath"`
}

func NewExtract() *Extract {
	return &Extract{}
}

type Transform struct {
	Delimiter string
	Regex     string
	JsonQuery string
}

func NewTransform() *Transform {
	return &Transform{}
}

type ContentExtractionEnumeration struct {
	FullPath    string
	Filename    string
	Content     string
	Lines       string
	LineNumbers string
}

func NewContentExtractionEnumeration() *ContentExtractionEnumeration {
	return &ContentExtractionEnumeration{}
}
