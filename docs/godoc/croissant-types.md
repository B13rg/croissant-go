# croissant

```go
import "github.com/b13rg/croissant-go/pkg/croissant"
```

Croissant spec filetypes and relations.

## Index

- [Variables](<#variables>)
- [type ContentExtractionEnumeration](<#ContentExtractionEnumeration>)
  - [func NewContentExtractionEnumeration\(\) \*ContentExtractionEnumeration](<#NewContentExtractionEnumeration>)
- [type DataSet](<#DataSet>)
  - [func NewDataSet\(\) \*DataSet](<#NewDataSet>)
  - [func NewFileSet\(\) \*DataSet](<#NewFileSet>)
- [type DataSource](<#DataSource>)
  - [func NewDataSource\(\) \*DataSource](<#NewDataSource>)
- [type DataType](<#DataType>)
  - [func NewDataType\(\) \*DataType](<#NewDataType>)
- [type Extract](<#Extract>)
  - [func NewExtract\(\) \*Extract](<#NewExtract>)
- [type Field](<#Field>)
  - [func NewField\(\) \*Field](<#NewField>)
- [type FileObject](<#FileObject>)
  - [func NewFileObject\(\) \*FileObject](<#NewFileObject>)
- [type FileResource](<#FileResource>)
  - [func NewFileResource\(\) \*FileResource](<#NewFileResource>)
- [type FileSet](<#FileSet>)
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

<a name="ContentExtractionEnumeration"></a>
## type [ContentExtractionEnumeration](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L221-L227>)



```go
type ContentExtractionEnumeration struct {
    FullPath    string
    Filename    string
    Content     string
    Lines       string
    LineNumbers string
}
```

<a name="NewContentExtractionEnumeration"></a>
### func [NewContentExtractionEnumeration](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L229>)

```go
func NewContentExtractionEnumeration() *ContentExtractionEnumeration
```



<a name="DataSet"></a>
## type [DataSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L4-L28>)



```go
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
```

<a name="NewDataSet"></a>
### func [NewDataSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L30>)

```go
func NewDataSet() *DataSet
```



<a name="NewFileSet"></a>
### func [NewFileSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L110>)

```go
func NewFileSet() *DataSet
```



<a name="DataSource"></a>
## type [DataSource](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L177-L185>)



```go
type DataSource struct {
    NType      string      `json:"@type"`
    FileObject *FileObject `json:"fileObject"`
    FileSet    *FileSet    `json:"fileSet"`
    RecordSet  *RecordSet  `json:"recordSet"`
    Extract    Extract     `json:"extract"`
    Transform  Transform   `json:"transform"`
    Format     Format      `json:"format"`
}
```

<a name="NewDataSource"></a>
### func [NewDataSource](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L187>)

```go
func NewDataSource() *DataSource
```



<a name="DataType"></a>
## type [DataType](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L139-L142>)



```go
type DataType struct {
    // MIME type
    DataType string
}
```

<a name="NewDataType"></a>
### func [NewDataType](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L144>)

```go
func NewDataType() *DataType
```



<a name="Extract"></a>
## type [Extract](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L191-L195>)



```go
type Extract struct {
    FileProperty ContentExtractionEnumeration
    Column       string `json:"column"`
    JsonPath     string `json:"jsonPath"`
}
```

<a name="NewExtract"></a>
### func [NewExtract](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L197>)

```go
func NewExtract() *Extract
```



<a name="Field"></a>
## type [Field](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L150-L159>)



```go
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
```

<a name="NewField"></a>
### func [NewField](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L161>)

```go
func NewField() *Field
```



<a name="FileObject"></a>
## type [FileObject](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L89-L97>)



```go
type FileObject struct {
    NType          string       `json:"@type"`
    Name           string       `json:"sc:name"`
    ContentURL     string       `json:"sc:contentUrl"`
    ContentSize    string       `json:"sc:contentSize"`
    EncodingFormat string       `json:"sc:encodingFormat"`
    Sha256         string       `json:"sc:sha256,omitempty"`
    ContainedIn    FileResource `json:"containedIn,omitempty"`
}
```

<a name="NewFileObject"></a>
### func [NewFileObject](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L99>)

```go
func NewFileObject() *FileObject
```



<a name="FileResource"></a>
## type [FileResource](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L115-L118>)

Type used to group data resource objects together.

```go
type FileResource struct {
    FileObject *FileObject
    FileSet    *FileSet
}
```

<a name="NewFileResource"></a>
### func [NewFileResource](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L120>)

```go
func NewFileResource() *FileResource
```



<a name="FileSet"></a>
## type [FileSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L103-L108>)



```go
type FileSet struct {
    NType       string       `json:"@type"`
    ContainedIn FileResource `json:"containedIn"`
    Includes    string       `json:"includes,omitempty"`
    Excludes    string       `json:"excludes,omitempty"`
}
```

<a name="Format"></a>
## type [Format](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L148>)



```go
type Format struct{}
```

<a name="JSON"></a>
## type [JSON](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L124>)



```go
type JSON string
```

<a name="RecordSet"></a>
## type [RecordSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L126-L133>)



```go
type RecordSet struct {
    NType    string   `json:"@type"`
    Field    []Field  `json:"field"`
    Key      []string `json:"key,omitempty"`
    Data     []JSON   `json:"data,omitempty"`
    Examples []JSON   `json:"examples,omitempty"`
    Source   Source   `json:"-"`
}
```

<a name="NewRecordSet"></a>
### func [NewRecordSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L135>)

```go
func NewRecordSet() *RecordSet
```



<a name="Source"></a>
## type [Source](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L166-L171>)

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
### func [NewSource](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L173>)

```go
func NewSource() *Source
```



<a name="Split"></a>
## type [Split](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L211-L215>)



```go
type Split struct {
    TrainSplit      string
    TestSplit       string
    ValidationSplit string
}
```

<a name="NewSplit"></a>
### func [NewSplit](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L217>)

```go
func NewSplit() *Split
```



<a name="Transform"></a>
## type [Transform](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L201-L205>)



```go
type Transform struct {
    Delimiter string
    Regex     string
    JsonQuery string
}
```

<a name="NewTransform"></a>
### func [NewTransform](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L207>)

```go
func NewTransform() *Transform
```

