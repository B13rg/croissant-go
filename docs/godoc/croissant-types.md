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
  - [func \(ref \*ClassRefList\) UnmarshalJSON\(data \[\]byte\) error](<#ClassRefList.UnmarshalJSON>)
- [type ContentExtractionEnumeration](<#ContentExtractionEnumeration>)
  - [func NewContentExtractionEnumeration\(\) \*ContentExtractionEnumeration](<#NewContentExtractionEnumeration>)
- [type DataSet](<#DataSet>)
  - [func NewDataSet\(\) \*DataSet](<#NewDataSet>)
  - [func NewDataSetFromPath\(path string\) \(\*DataSet, error\)](<#NewDataSetFromPath>)
  - [func NewFileSet\(\) \*DataSet](<#NewFileSet>)
- [type DataSource](<#DataSource>)
  - [func NewDataSource\(\) \*DataSource](<#NewDataSource>)
- [type DataType](<#DataType>)
  - [func NewDataType\(\) \*DataType](<#NewDataType>)
- [type Distribution](<#Distribution>)
  - [func \(d \*Distribution\) UnmarshalJSON\(data \[\]byte\) error](<#Distribution.UnmarshalJSON>)
- [type DistributionItem](<#DistributionItem>)
- [type Extract](<#Extract>)
  - [func NewExtract\(\) \*Extract](<#NewExtract>)
- [type Field](<#Field>)
  - [func NewField\(\) \*Field](<#NewField>)
- [type FieldRef](<#FieldRef>)
- [type FileObject](<#FileObject>)
  - [func NewFileObject\(\) \*FileObject](<#NewFileObject>)
  - [func \(\*FileObject\) Update\(\) error](<#FileObject.Update>)
  - [func \(\*FileObject\) Validate\(\) error](<#FileObject.Validate>)
  - [func \(\*FileObject\) ValidateHash\(\) error](<#FileObject.ValidateHash>)
- [type FileSet](<#FileSet>)
  - [func \(\*FileSet\) Update\(\) error](<#FileSet.Update>)
  - [func \(\*FileSet\) Validate\(\) error](<#FileSet.Validate>)
- [type Format](<#Format>)
- [type JSON](<#JSON>)
- [type RecordSet](<#RecordSet>)
  - [func NewRecordSet\(\) \*RecordSet](<#NewRecordSet>)
- [type Source](<#Source>)
  - [func NewSource\(\) \*Source](<#NewSource>)
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
## type [ClassRefItem](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/references.go#L8-L10>)



```go
type ClassRefItem struct {
    Id string `json="@id"`
}
```

<a name="ClassRefList"></a>
## type [ClassRefList](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/references.go#L12-L14>)



```go
type ClassRefList struct {
    Items []ClassRefItem
}
```

<a name="ClassRefList.UnmarshalJSON"></a>
### func \(\*ClassRefList\) [UnmarshalJSON](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/references.go#L16>)

```go
func (ref *ClassRefList) UnmarshalJSON(data []byte) error
```



<a name="ContentExtractionEnumeration"></a>
## type [ContentExtractionEnumeration](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/datasource.go#L52-L63>)



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
### func [NewContentExtractionEnumeration](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/datasource.go#L65>)

```go
func NewContentExtractionEnumeration() *ContentExtractionEnumeration
```



<a name="DataSet"></a>
## type [DataSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/dataset.go#L14-L68>)

\[Dataset Class\]\(https://docs.mlcommons.org/croissant/docs/croissant-spec.html#dataset-level-information\) Based on https://docs.mlcommons.org/croissant/docs/croissant-spec.html#schemaorgdataset

```go
type DataSet struct {

    // Context alias definitions to make rest of document shorter.
    Context map[string]interface{} `json:"@context"`
    // Must be `schema.org/Dataset`.
    NType string `json:"@type"`
    // Schema version the croissant file conforms to.
    ConformsTo string `json:"dct:conformsTo"`
    // Description of the dataset
    Description string `json:"description"`
    // Licenses of the dataset.
    // Spec suggests using references from https://spdx.org/licenses/.
    License types.StringOrSlice `json:"license"`
    // The name of the dataset
    Name string `json:"name"`
    // Url of the dataset, usually a webpage.
    URL string `json:"url"`
    // One or more Person or Organizations that created the dataset.
    Creator []string `json:"creator"`
    // The date the dataset was published.
    DatePublished string `json:"datePublished"`

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
    // License that applies the the croissant metadata.
    SdLicense []string `json:"sdLicense,omitempty"`
    // Language of the content of the dataset.
    InLanguage []string `json:"inLanguage,omitempty"`

    // List of FileObjects and FileSets associated with the dataset.
    // Modified from schema.org/Dataset.
    // Required.
    Distribution Distribution `json:"distribution,omitempty"`
    //Whether the dataset is a live dataset (in-process of being updated).
    IsLiveDataset bool `json:"isLiveDataset,omitempty"`
    // A citation to the dataset itself.
    CiteAs string `json:"citeAs,omitempty"`
}
```

<a name="NewDataSet"></a>
### func [NewDataSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/dataset.go#L70>)

```go
func NewDataSet() *DataSet
```



<a name="NewDataSetFromPath"></a>
### func [NewDataSetFromPath](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/dataset.go#L80>)

```go
func NewDataSetFromPath(path string) (*DataSet, error)
```



<a name="NewFileSet"></a>
### func [NewFileSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L50>)

```go
func NewFileSet() *DataSet
```



<a name="DataSource"></a>
## type [DataSource](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/datasource.go#L4-L18>)



```go
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
```

<a name="NewDataSource"></a>
### func [NewDataSource](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/datasource.go#L20>)

```go
func NewDataSource() *DataSource
```



<a name="DataType"></a>
## type [DataType](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/recordset.go#L25-L28>)



```go
type DataType struct {
    // MIME type
    DataType string
}
```

<a name="NewDataType"></a>
### func [NewDataType](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/recordset.go#L30>)

```go
func NewDataType() *DataType
```



<a name="Distribution"></a>
## type [Distribution](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/dataset.go#L93-L95>)



```go
type Distribution struct {
    Items []DistributionItem
}
```

<a name="Distribution.UnmarshalJSON"></a>
### func \(\*Distribution\) [UnmarshalJSON](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/dataset.go#L97>)

```go
func (d *Distribution) UnmarshalJSON(data []byte) error
```



<a name="DistributionItem"></a>
## type [DistributionItem](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/dataset.go#L91>)

Type used to group data resource objects together.

```go
type DistributionItem interface{}
```

<a name="Extract"></a>
## type [Extract](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/datasource.go#L26-L33>)



```go
type Extract struct {
    // Extraction method.
    FileProperty ContentExtractionEnumeration
    // Name of the column (field) that contains values.
    Column string `json:"column"`
    // A JSON path expression that obtains values.
    JsonPath string `json:"jsonPath"`
}
```

<a name="NewExtract"></a>
### func [NewExtract](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/datasource.go#L35>)

```go
func NewExtract() *Extract
```



<a name="Field"></a>
## type [Field](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/field.go#L4-L20>)



```go
type Field struct {
    NType string `json:"@type"`
    // The source of data for the field.
    Source Source `json:"source"`
    // The data types that correspond to the Field.
    DataType []*DataType `json:"dataType"`
    // If true the Field is a list of DataType values.
    Repeated bool `json:"repeated,omitempty"`
    // A property URL that is equivalent to this field
    EquivalentProperty string `json:"equivalentProperty,omitempty"`
    // References to other fields in a different RecordSet.
    References ClassRefList `json:"references,omitempty"`
    // List of Fields nested inside this one.
    SubField []Field `json:"subField,omitempty"`
    // References to other Fields in the same RecordSet.
    ParentField ClassRefList `json:"parentField,omitempty"`
}
```

<a name="NewField"></a>
### func [NewField](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/field.go#L22>)

```go
func NewField() *Field
```



<a name="FieldRef"></a>
## type [FieldRef](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/field.go#L26-L28>)



```go
type FieldRef struct {
    Field ClassRefItem `json="field"`
}
```

<a name="FileObject"></a>
## type [FileObject](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L4-L20>)



```go
type FileObject struct {
    // Must be FileObject
    NType string `json:"@type"`
    // The name of the file.
    Name string `json:"sc:name"`
    // URL to the actual bytes of the file object.
    ContentURL string `json:"sc:contentUrl"`
    // File size in [mega/kilo/...]bytes.
    // Defaults to bytes if unit not specified.
    ContentSize string `json:"sc:contentSize"`
    // Format of the file given as a MIME type.
    EncodingFormat string `json:"sc:encodingFormat"`
    // Checksum of the file contents.
    Sha256 string `json:"sc:sha256,omitempty"`
    // Another FileObject or FileSet this resource is contained in.
    ContainedIn ClassRefList `json:"containedIn,omitempty"`
}
```

<a name="NewFileObject"></a>
### func [NewFileObject](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L22>)

```go
func NewFileObject() *FileObject
```



<a name="FileObject.Update"></a>
### func \(\*FileObject\) [Update](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L35>)

```go
func (*FileObject) Update() error
```

Update FileObject struct from resource

<a name="FileObject.Validate"></a>
### func \(\*FileObject\) [Validate](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L26>)

```go
func (*FileObject) Validate() error
```



<a name="FileObject.ValidateHash"></a>
### func \(\*FileObject\) [ValidateHash](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L30>)

```go
func (*FileObject) ValidateHash() error
```



<a name="FileSet"></a>
## type [FileSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L39-L48>)



```go
type FileSet struct {
    // Must be FileSet.
    NType string `json:"@type"`
    // The FileSet or FileObject the resource is contained in.
    ContainedIn ClassRefList `json:"containedIn"`
    // A glob pattern of files to include.
    Includes string `json:"includes,omitempty"`
    // A glob patter of files to exclude.
    Excludes string `json:"excludes,omitempty"`
}
```

<a name="FileSet.Update"></a>
### func \(\*FileSet\) [Update](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L59>)

```go
func (*FileSet) Update() error
```

Update FileSet struct from resource

<a name="FileSet.Validate"></a>
### func \(\*FileSet\) [Validate](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/file_resources.go#L54>)

```go
func (*FileSet) Validate() error
```



<a name="Format"></a>
## type [Format](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/datasource.go#L24>)



```go
type Format struct{}
```

<a name="JSON"></a>
## type [JSON](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/recordset.go#L4>)



```go
type JSON string
```

<a name="RecordSet"></a>
## type [RecordSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/recordset.go#L6-L19>)



```go
type RecordSet struct {
    // Must be RecordSet
    NType string `json:"@type"`
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
```

<a name="NewRecordSet"></a>
### func [NewRecordSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/recordset.go#L21>)

```go
func NewRecordSet() *RecordSet
```



<a name="Source"></a>
## type [Source](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/field.go#L31-L36>)

Type used to group data sources.

```go
type Source struct {
    DataSource *DataSource
    FileObject *FileObject
    FileSet    *FileSet
    RecordSet  *RecordSet
}
```

<a name="NewSource"></a>
### func [NewSource](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/field.go#L38>)

```go
func NewSource() *Source
```



<a name="Split"></a>
## type [Split](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/recordset.go#L34-L38>)



```go
type Split struct {
    TrainSplit      string
    TestSplit       string
    ValidationSplit string
}
```

<a name="NewSplit"></a>
### func [NewSplit](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/recordset.go#L40>)

```go
func NewSplit() *Split
```



<a name="Transform"></a>
## type [Transform](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/datasource.go#L39-L46>)



```go
type Transform struct {
    // Split data source string on character.
    Delimiter string
    // Apply regex to data source.
    Regex string
    // A JSON query to evaluate against the data source.
    JsonQuery string
}
```

<a name="NewTransform"></a>
### func [NewTransform](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/datasource.go#L48>)

```go
func NewTransform() *Transform
```

