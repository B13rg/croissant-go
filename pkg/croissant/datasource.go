// Croissant spec filetypes and relations.
package croissant

type DataSource struct {
	// Must be DataSource
	NType string `json:"@type,omitempty"`
	// Node ID
	NId string `json:"@id,omitempty"`
	// The name of the referenced FileObject source of the data.
	FileObject ClassRefItem `json:"fileObject,omitempty"`
	// The name of the reference RecordSet source.
	FileSet ClassRefItem `json:"fileSet,omitempty"`
	// The name of the referenced RecordSet source.
	RecordSet ClassRefItem `json:"recordSet,omitempty"`
	// The extraction method from the provided source.
	Extract Extract `json:"extract,omitempty"`
	// Transformations to apply to data after extraction.
	Transform Transform `json:"transform,omitempty"`
	// A format to parse data values from text.
	Format Format `json:"format,omitempty"`
}

func NewDataSource() *DataSource {
	return &DataSource{}
}

type Format struct{}

type Extract struct {
	// Extraction method.
	FileProperty string `json:"fileProperty,omitempty"`
	// Name of the column (field) that contains values.
	Column string `json:"column,omitempty"`
	// A JSON path expression that obtains values.
	JsonPath string `json:"jsonPath,omitempty"`
}

func NewExtract() *Extract {
	return &Extract{}
}

type Transform struct {
	// Split data source string on character.
	Delimiter string `json:"delimiter,omitempty"`
	// Apply regex to data source.
	Regex string `json:"regex,omitempty"`
	// the path to extract json from.
	JsonPath string `json:"jsonPath,omitempty"`
	// A JSON query to evaluate against the data source.
	JsonQuery string `json:"jsonquery,omitempty"`
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
