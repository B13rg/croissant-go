// Croissant spec filetypes and relations.
package croissant

type DataSet struct {
	Context    map[string]interface{} `json:"@context"`
	NType      string                 `json:"@type"`
	ConformsTo string                 `json:"dct:conformsTo"`

	Description   string   `json:"description"`
	License       []string `json:"license"`
	Name          string   `json:"name"`
	URL           string   `json:"url"`
	Creator       []string `json:"creator"`
	DatePublished string   `json:"datePublished"`

	Keywords     []string `json:"keywords,omitempty"`
	Publisher    []string `json:"publisher,omitempty"`
	Version      string   `json:"version,omitempty"`
	DateCreated  string   `json:"dateCreated,omitempty"`
	DateModified string   `json:"dateModified,omitempty"`
	SameAs       []string `json:"sameAs,omitempty"`
	SdLicense    []string `json:"sdLicense,omitempty"`
	InLanguage   []string `json:"inLanguage,omitempty"`

	Distribution  []FileResource `json:"distribution,omitempty"`
	IsLiveDataset bool           `json:"isLiveDataset,omitempty"`
	CiteAs        string         `json:"citeAs,omitempty"`
}

func NewDataSet() *DataSet {
	return &DataSet{
		Context:     SuggestedContext,
		NType:       "sc:DataSet",
		ConformsTo:  "http://mlcommons.org/croissant/1.0",
		Description: "",
		License:     []string{},
	}
}

var SuggestedContext = map[string]interface{}{
	"@language":  "en",
	"@vocab":     "https://schema.org/",
	"sc":         "https://schema.org/",
	"cr":         "http://mlcommons.org/croissant/",
	"rai":        "http://mlcommons.org/croissant/RAI/",
	"dct":        "http://purl.org/dc/terms/",
	"citeAs":     "cr:citeAs",
	"column":     "cr:column",
	"conformsTo": "dct:conformsTo",
	"data": map[string]interface{}{
		"@id":   "cr:data",
		"@type": "@json",
	},
	"dataType": map[string]interface{}{
		"@id":   "cr:dataType",
		"@type": "@vocab",
	},
	"examples": map[string]interface{}{
		"@id":   "cr:examples",
		"@type": "@json",
	},
	"extract":       "cr:extract",
	"field":         "cr:field",
	"fileProperty":  "cr:fileProperty",
	"fileObject":    "cr:fileObject",
	"fileSet":       "cr:fileSet",
	"format":        "cr:format",
	"includes":      "cr:includes",
	"isLiveDataset": "cr:isLiveDataset",
	"jsonPath":      "cr:jsonPath",
	"key":           "cr:key",
	"md5":           "cr:md5",
	"parentField":   "cr:parentField",
	"path":          "cr:path",
	"recordSet":     "cr:recordSet",
	"references":    "cr:references",
	"regex":         "cr:regex",
	"repeated":      "cr:repeated",
	"replace":       "cr:replace",
	"separator":     "cr:separator",
	"source":        "cr:source",
	"subField":      "cr:subField",
	"transform":     "cr:transform",
}

type FileObject struct {
	NType          string       `json:"@type"`
	Name           string       `json:"sc:name"`
	ContentURL     string       `json:"sc:contentUrl"`
	ContentSize    string       `json:"sc:contentSize"`
	EncodingFormat string       `json:"sc:encodingFormat"`
	Sha256         string       `json:"sc:sha256,omitempty"`
	ContainedIn    FileResource `json:"containedIn,omitempty"`
}

func NewFileObject() *FileObject {
	return &FileObject{}
}

type FileSet struct {
	NType       string       `json:"@type"`
	ContainedIn FileResource `json:"containedIn"`
	Includes    string       `json:"includes,omitempty"`
	Excludes    string       `json:"excludes,omitempty"`
}

func NewFileSet() *DataSet {
	return &DataSet{}
}

// Type used to group data resource objects together.
type FileResource struct {
	FileObject *FileObject
	FileSet    *FileSet
}

func NewFileResource() *FileResource {
	return &FileResource{}
}

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

type Format struct{}

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

type Split struct {
	TrainSplit      string
	TestSplit       string
	ValidationSplit string
}

func NewSplit() *Split {
	return &Split{}
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
