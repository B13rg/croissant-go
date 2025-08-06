// Croissant spec filetypes and relations.
package croissant

type DataSet struct {
	Description   string
	License       []string
	Name          string
	URL           string
	Creator       []string
	DatePublished string

	Keywords     []string
	Publisher    []string
	Version      string
	DateCreated  string
	DateModified string
	SameAs       []string
	SdLicense    []string
	InLanguage   []string

	Distribution  []FileResource
	isLiveDataset bool
	citeAs        string
}

type FileObject struct {
	ContainedIn FileResource
}

type FileSet struct {
	ContainedIn FileResource
	Includes    string
	Excludes    string
}

type FileResource struct {
	FileObject *FileObject
	FileSet    *FileSet
}

type JSON string

type RecordSet struct {
	Key      []*Field
	Data     []JSON
	Examples []JSON
	Source   Source
}

type DataType struct {
	// MIME type
	DataType string
}

type Format struct{}

type Field struct {
	Source             Source
	Repeated           bool
	References         []*Field
	SubField           []*Field
	ParentField        []*Field
	EquivalentProperty string
}

type Source struct {
	DataSource *DataSource
	FileObject *FileObject
	FileSet    *FileSet
	RecordSet  *RecordSet
}

type DataSource struct {
	FileObject *FileObject
	FileSet    *FileSet
	RecordSet  *RecordSet
	Extract    Extract
	Transform  Transform
	Format     Format
}

type Extract struct {
	FileProperty ContentExtractionEnumeration
	Column       string
	JsonPath     string
}

type Transform struct {
	Delimiter string
	Regex     string
	JsonQuery string
}

type Split struct {
	TrainSplit      string
	TestSplit       string
	ValidationSplit string
}

type ContentExtractionEnumeration struct {
	FullPath string
	Filename string
	Content string
	Lines string
	LineNumbers string
}
