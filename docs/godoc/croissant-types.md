# croissant

```go
import "github.com/b13rg/croissant-go/pkg/croissant"
```

Croissant spec filetypes and relations.

Croissant spec filetypes and relations.

Croissant spec filetypes and relations.

Croissant spec filetypes and relations.

Croissant spec filetypes and relations.

## Index

- [Variables](<#variables>)
- [type ClassRefItem](<#ClassRefItem>)
- [type ClassRefList](<#ClassRefList>)
  - [func \(ref ClassRefList\) MarshalJSON\(\) \(\[\]byte, error\)](<#ClassRefList.MarshalJSON>)
  - [func \(ref \*ClassRefList\) UnmarshalJSON\(data \[\]byte\) error](<#ClassRefList.UnmarshalJSON>)
- [type ContentExtractionEnumeration](<#ContentExtractionEnumeration>)
  - [func NewContentExtractionEnumeration\(\) \*ContentExtractionEnumeration](<#NewContentExtractionEnumeration>)
- [type DataSet](<#DataSet>)
  - [func NewDataSet\(\) \*DataSet](<#NewDataSet>)
  - [func NewDataSetFromPath\(filePath string\) \(\*DataSet, error\)](<#NewDataSetFromPath>)
  - [func NewFileSet\(\) \*DataSet](<#NewFileSet>)
  - [func \(ds \*DataSet\) Validate\(\) \(\[\]types.CroissantWarning, \[\]types.CroissantError\)](<#DataSet.Validate>)
  - [func \(ds \*DataSet\) WriteToFile\(path string\) error](<#DataSet.WriteToFile>)
- [type DataSource](<#DataSource>)
  - [func NewDataSource\(\) \*DataSource](<#NewDataSource>)
- [type DataType](<#DataType>)
  - [func NewDataType\(mimeType string\) \*DataType](<#NewDataType>)
- [type Distribution](<#Distribution>)
  - [func \(d \*Distribution\) UnmarshalJSON\(data \[\]byte\) error](<#Distribution.UnmarshalJSON>)
- [type DistributionItem](<#DistributionItem>)
- [type Extract](<#Extract>)
  - [func NewExtract\(\) \*Extract](<#NewExtract>)
- [type Field](<#Field>)
  - [func NewField\(\) \*Field](<#NewField>)
  - [func \(obj \*Field\) Validate\(\) \(\[\]types.CroissantWarning, \[\]types.CroissantError\)](<#Field.Validate>)
- [type FieldRef](<#FieldRef>)
- [type FieldRefSlice](<#FieldRefSlice>)
  - [func \(ref FieldRefSlice\) MarshalJSON\(\) \(\[\]byte, error\)](<#FieldRefSlice.MarshalJSON>)
  - [func \(ref \*FieldRefSlice\) UnmarshalJSON\(data \[\]byte\) error](<#FieldRefSlice.UnmarshalJSON>)
- [type FileObject](<#FileObject>)
  - [func NewFileObject\(\) \*FileObject](<#NewFileObject>)
  - [func \(\*FileObject\) Update\(\) error](<#FileObject.Update>)
  - [func \(obj \*FileObject\) Validate\(\) \(\[\]types.CroissantWarning, \[\]types.CroissantError\)](<#FileObject.Validate>)
  - [func \(\*FileObject\) ValidateHash\(\) error](<#FileObject.ValidateHash>)
- [type FileSet](<#FileSet>)
  - [func \(\*FileSet\) Update\(\) error](<#FileSet.Update>)
  - [func \(obj \*FileSet\) Validate\(\) \(\[\]types.CroissantWarning, \[\]types.CroissantError\)](<#FileSet.Validate>)
- [type Format](<#Format>)
- [type RecordSet](<#RecordSet>)
  - [func NewRecordSet\(\) \*RecordSet](<#NewRecordSet>)
  - [func \(obj \*RecordSet\) Validate\(\) \(\[\]types.CroissantWarning, \[\]types.CroissantError\)](<#RecordSet.Validate>)
- [type Split](<#Split>)
  - [func NewSplit\(\) \*Split](<#NewSplit>)
- [type Transform](<#Transform>)
  - [func NewTransform\(\) \*Transform](<#NewTransform>)


## Variables

<a name="SuggestedContext"></a>The suggested context to use in a Croissant Json\-LD file.

```go
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
```

<a name="ClassRefItem"></a>
## type [ClassRefItem](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/references.go#L9-L12>)



```go
type ClassRefItem struct {
    // ID of the resource.
    ID string `json:"@id,omitempty"`
}
```

<a name="ClassRefList"></a>
## type [ClassRefList](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/references.go#L14>)



```go
type ClassRefList []ClassRefItem
```

<a name="ClassRefList.MarshalJSON"></a>
### func \(ClassRefList\) [MarshalJSON](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/references.go#L38>)

```go
func (ref ClassRefList) MarshalJSON() ([]byte, error)
```



<a name="ClassRefList.UnmarshalJSON"></a>
### func \(\*ClassRefList\) [UnmarshalJSON](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/references.go#L16>)

```go
func (ref *ClassRefList) UnmarshalJSON(data []byte) error
```



<a name="ContentExtractionEnumeration"></a>
## type [ContentExtractionEnumeration](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/datasource.go#L57-L68>)



```go
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
```

<a name="NewContentExtractionEnumeration"></a>
### func [NewContentExtractionEnumeration](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/datasource.go#L70>)

```go
func NewContentExtractionEnumeration() *ContentExtractionEnumeration
```



<a name="DataSet"></a>
## type [DataSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/dataset.go#L14-L70>)

\[Dataset Class\]\(https://docs.mlcommons.org/croissant/docs/croissant-spec.html#dataset-level-information\) Based on https://docs.mlcommons.org/croissant/docs/croissant-spec.html#schemaorgdataset

```go
type DataSet struct {

    // Context alias definitions to make rest of document shorter.
    Context map[string]interface{} `json:"@context"`
    // Must be `schema.org/Dataset`.
    Type string `json:"@type"`
    // The name of the dataset
    Name string `json:"name"`
    // Description of the dataset
    Description string `json:"description"`
    // Schema version the croissant file conforms to.
    ConformsTo string `json:"conformsTo"`
    // A citation to the dataset itself.
    CiteAs string `json:"citeAs,omitempty"`
    // Licenses of the dataset.
    // Spec suggests using references from https://spdx.org/licenses/.
    License types.StringOrSlice `json:"license"`
    // Url of the dataset, usually a webpage.
    URL string `json:"url"`
    // One or more Person or Organizations that created the dataset.
    Creator []string `json:"creator,omitempty"`
    // The date the dataset was published.
    DatePublished string `json:"datePublished,omitempty"`

    // Keywords associated with the text
    Keywords []string `json:"keywords,omitempty"`
    // Publisher of the dataset, sometimes distinct from creator.
    Publisher []string `json:"publisher,omitempty"`
    // Version of the dataset.
    // Either an single int, or a MAJOR.MINOR.PATCH sematic version string.
    // [Semantic Versioning 2.0.0](https://semver.org/spec/v2.0.0.html)
    Version string `json:"version,omitempty"`
    // Date the dataset was initially created
    DateCreated string `json:"dateCreated,omitempty"`
    // Date the dataset was last modified
    DateModified string `json:"dateModified,omitempty"`
    // List of URLs that represent the same dataset as this one.
    SameAs []string `json:"sameAs,omitempty"`
    // License that applies to the croissant metadata.
    SdLicense []string `json:"sdLicense,omitempty"`
    // Language of the content of the dataset.
    InLanguage []string `json:"inLanguage,omitempty"`

    // Whether the dataset is a live dataset (in-process of being updated).
    IsLiveDataset bool `json:"isLiveDataset,omitempty"`

    // List of FileObjects and FileSets associated with the dataset.
    // Modified from schema.org/Dataset.
    // Required.
    Distribution Distribution `json:"distribution,omitempty"`

    // List of RecordSets associated with the dataset
    RecordSets []RecordSet `json:"recordSet"`
}
```

<a name="NewDataSet"></a>
### func [NewDataSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/dataset.go#L72>)

```go
func NewDataSet() *DataSet
```



<a name="NewDataSetFromPath"></a>
### func [NewDataSetFromPath](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/dataset.go#L235>)

```go
func NewDataSetFromPath(filePath string) (*DataSet, error)
```



<a name="NewFileSet"></a>
### func [NewFileSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L173>)

```go
func NewFileSet() *DataSet
```



<a name="DataSet.Validate"></a>
### func \(\*DataSet\) [Validate](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/dataset.go#L82>)

```go
func (ds *DataSet) Validate() ([]types.CroissantWarning, []types.CroissantError)
```



<a name="DataSet.WriteToFile"></a>
### func \(\*DataSet\) [WriteToFile](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/dataset.go#L245>)

```go
func (ds *DataSet) WriteToFile(path string) error
```



<a name="DataSource"></a>
## type [DataSource](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/datasource.go#L4-L21>)



```go
type DataSource struct {
    // Must be DataSource
    NType *string `json:"@type,omitempty"`
    // Node ID
    ClassRefItem
    // The name of the referenced FileObject source of the data.
    FileObject *ClassRefItem `json:"fileObject,omitempty"`
    // The name of the reference RecordSet source.
    FileSet *ClassRefItem `json:"fileSet,omitempty"`
    // The name of the referenced RecordSet source.
    RecordSet *ClassRefItem `json:"recordSet,omitempty"`
    // The extraction method from the provided source.
    Extract *Extract `json:"extract,omitempty"`
    // Transformations to apply to data after extraction.
    Transform *Transform `json:"transform,omitempty"`
    // A format to parse data values from text.
    Format *Format `json:"format,omitempty"`
}
```

<a name="NewDataSource"></a>
### func [NewDataSource](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/datasource.go#L23>)

```go
func NewDataSource() *DataSource
```



<a name="DataType"></a>
## type [DataType](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/recordset.go#L63-L66>)



```go
type DataType struct {
    // MIME type
    DataType string
}
```

<a name="NewDataType"></a>
### func [NewDataType](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/recordset.go#L68>)

```go
func NewDataType(mimeType string) *DataType
```



<a name="Distribution"></a>
## type [Distribution](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L11>)

Type used to group data resource objects together.

```go
type Distribution []DistributionItem
```

<a name="Distribution.UnmarshalJSON"></a>
### func \(\*Distribution\) [UnmarshalJSON](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L16>)

```go
func (d *Distribution) UnmarshalJSON(data []byte) error
```



<a name="DistributionItem"></a>
## type [DistributionItem](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L12-L14>)



```go
type DistributionItem interface {
    Validate() ([]types.CroissantWarning, []types.CroissantError)
}
```

<a name="Extract"></a>
## type [Extract](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/datasource.go#L29-L36>)



```go
type Extract struct {
    // Extraction method.
    FileProperty string `json:"fileProperty,omitempty"`
    // Name of the column (field) that contains values.
    Column string `json:"column,omitempty"`
    // A JSON path expression that obtains values.
    JsonPath string `json:"jsonPath,omitempty"`
}
```

<a name="NewExtract"></a>
### func [NewExtract](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/datasource.go#L38>)

```go
func NewExtract() *Extract
```



<a name="Field"></a>
## type [Field](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/field.go#L10-L33>)



```go
type Field struct {
    // Must be field.
    Type string `json:"@type"`
    // Node ID.
    ClassRefItem
    // Name of the Field.
    Name string `json:"name"`
    // Description of the Field.
    Description string `json:"description"`
    // The data types that correspond to the Field.
    DataType types.StringOrSlice `json:"dataType,omitempty"`
    // The source of data for the field.
    Source *DataSource `json:"source,omitempty"`
    // If true the Field is a list of DataType values.
    Repeated bool `json:"repeated,omitempty"`
    // A property URL that is equivalent to this field
    EquivalentProperty string `json:"equivalentProperty,omitempty"`
    // References to other fields in a different RecordSet.
    References FieldRefSlice `json:"references,omitempty"`
    // List of Fields nested inside this one.
    SubField []Field `json:"subField,omitempty"`
    // References to other Fields in the same RecordSet.
    ParentField FieldRefSlice `json:"parentField,omitempty"`
}
```

<a name="NewField"></a>
### func [NewField](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/field.go#L35>)

```go
func NewField() *Field
```



<a name="Field.Validate"></a>
### func \(\*Field\) [Validate](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/field.go#L39>)

```go
func (obj *Field) Validate() ([]types.CroissantWarning, []types.CroissantError)
```



<a name="FieldRef"></a>
## type [FieldRef](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/field.go#L86-L88>)



```go
type FieldRef struct {
    Field ClassRefItem `json:"field"`
}
```

<a name="FieldRefSlice"></a>
## type [FieldRefSlice](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/field.go#L90>)



```go
type FieldRefSlice []FieldRef
```

<a name="FieldRefSlice.MarshalJSON"></a>
### func \(FieldRefSlice\) [MarshalJSON](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/field.go#L114>)

```go
func (ref FieldRefSlice) MarshalJSON() ([]byte, error)
```



<a name="FieldRefSlice.UnmarshalJSON"></a>
### func \(\*FieldRefSlice\) [UnmarshalJSON](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/field.go#L92>)

```go
func (ref *FieldRefSlice) UnmarshalJSON(data []byte) error
```



<a name="FileObject"></a>
## type [FileObject](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L67-L87>)



```go
type FileObject struct {
    // Must be FileObject
    Type string `json:"@type"`
    // Node ID.
    ClassRefItem
    // The name of the file.
    Name string `json:"name"`
    // Description of file.
    Description string `json:"description"`
    // URL to the actual bytes of the file object.
    ContentURL string `json:"contentUrl"`
    // File size in [mega/kilo/...]bytes.
    // Defaults to bytes if unit not specified.
    ContentSize string `json:"contentSize,omitempty"`
    // Format of the file given as a MIME type.
    EncodingFormat string `json:"encodingFormat"`
    // Checksum of the file contents.
    Sha256 string `json:"sha256,omitempty"`
    // Another FileObject or FileSet this resource is contained in.
    ContainedIn ClassRefList `json:"containedIn,omitempty"`
}
```

<a name="NewFileObject"></a>
### func [NewFileObject](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L89>)

```go
func NewFileObject() *FileObject
```



<a name="FileObject.Update"></a>
### func \(\*FileObject\) [Update](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L150>)

```go
func (*FileObject) Update() error
```

Update FileObject struct from resource.

<a name="FileObject.Validate"></a>
### func \(\*FileObject\) [Validate](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L93>)

```go
func (obj *FileObject) Validate() ([]types.CroissantWarning, []types.CroissantError)
```



<a name="FileObject.ValidateHash"></a>
### func \(\*FileObject\) [ValidateHash](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L145>)

```go
func (*FileObject) ValidateHash() error
```



<a name="FileSet"></a>
## type [FileSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L154-L171>)



```go
type FileSet struct {
    // Must be FileSet.
    Type string `json:"@type"`
    // Node ID
    ClassRefItem
    // Name of FileSet
    Name string `json:"name"`
    // Description of FileSet
    Description string `json:"description"`
    // The FileSet or FileObject the resource is contained in.
    ContainedIn ClassRefList `json:"containedIn"`
    // MIME type
    EncodingFormat string `json:"encodingFormat"`
    // A glob pattern of files to include.
    Includes string `json:"includes"`
    // A glob patter of files to exclude.
    Excludes string `json:"excludes,omitempty"`
}
```

<a name="FileSet.Update"></a>
### func \(\*FileSet\) [Update](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L208>)

```go
func (*FileSet) Update() error
```

Update FileSet struct from resource.

<a name="FileSet.Validate"></a>
### func \(\*FileSet\) [Validate](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L177>)

```go
func (obj *FileSet) Validate() ([]types.CroissantWarning, []types.CroissantError)
```



<a name="Format"></a>
## type [Format](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/datasource.go#L27>)



```go
type Format struct{}
```

<a name="RecordSet"></a>
## type [RecordSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/recordset.go#L6-L25>)



```go
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
```

<a name="NewRecordSet"></a>
### func [NewRecordSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/recordset.go#L27>)

```go
func NewRecordSet() *RecordSet
```



<a name="RecordSet.Validate"></a>
### func \(\*RecordSet\) [Validate](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/recordset.go#L31>)

```go
func (obj *RecordSet) Validate() ([]types.CroissantWarning, []types.CroissantError)
```



<a name="Split"></a>
## type [Split](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/recordset.go#L72-L76>)



```go
type Split struct {
    TrainSplit      string
    TestSplit       string
    ValidationSplit string
}
```

<a name="NewSplit"></a>
### func [NewSplit](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/recordset.go#L78>)

```go
func NewSplit() *Split
```



<a name="Transform"></a>
## type [Transform](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/datasource.go#L42-L51>)



```go
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
```

<a name="NewTransform"></a>
### func [NewTransform](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/datasource.go#L53>)

```go
func NewTransform() *Transform
```

