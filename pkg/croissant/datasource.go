// Croissant spec filetypes and relations.
package croissant

type DataSource struct {
	NType string `json:"@type"`
	// The name of the referenced FileObject source of the data.
	FileObject *FileObject `json:"fileObject"`
	// The name of the reference RecordSet source.
	FileSet *FileSet `json:"fileSet"`
	// The name of the referenced RecordSet source.
	RecordSet *RecordSet `json:"recordSet"`
	// The extraction method from the provided source.
	Extract Extract `json:"extract"`
	// Transformations to apply to data after extraction.
	Transform Transform `json:"transform"`
	// A format to parse data values from text.
	Format Format `json:"format"`
}

func NewDataSource() *DataSource {
	return &DataSource{}
}

type Format struct{}

type Extract struct {
	// Extraction method.
	FileProperty ContentExtractionEnumeration
	// Name of the column (field) that contains values.
	Column string `json:"column"`
	// A JSON path expression that obtains values.
	JsonPath string `json:"jsonPath"`
}

func NewExtract() *Extract {
	return &Extract{}
}

type Transform struct {
	// Split data source string on character.
	Delimiter string
	// Apply regex to data source.
	Regex string
	// A JSON query to evaluate against the data source.
	JsonQuery string
}

func NewTransform() *Transform {
	return &Transform{}
}

type ContentExtractionEnumeration struct {
	// Full path to file, from Croissant extraction or download folders.
	FullPath string
	// Name of the file (no path).
	Filename string
	// Byte content of the file.
	Content string
	// Byte content of each line of the file.
	Lines string
	// The numbers of each line in the file.
	LineNumbers string
}

func NewContentExtractionEnumeration() *ContentExtractionEnumeration {
	return &ContentExtractionEnumeration{}
}
