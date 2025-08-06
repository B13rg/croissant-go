# croissant

```go
import "github.com/b13rg/croissant-go/pkg/croissant"
```

Croissant spec filetypes and relations.

## Index

- [type ContentExtractionEnumeration](<#ContentExtractionEnumeration>)
- [type DataSet](<#DataSet>)
- [type DataSource](<#DataSource>)
- [type DataType](<#DataType>)
- [type Extract](<#Extract>)
- [type Field](<#Field>)
- [type FileObject](<#FileObject>)
- [type FileResource](<#FileResource>)
- [type FileSet](<#FileSet>)
- [type Format](<#Format>)
- [type JSON](<#JSON>)
- [type RecordSet](<#RecordSet>)
- [type Source](<#Source>)
- [type Split](<#Split>)
- [type Transform](<#Transform>)


<a name="ContentExtractionEnumeration"></a>
## type [ContentExtractionEnumeration](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L100-L106>)



```go
type ContentExtractionEnumeration struct {
    FullPath    string
    Filename    string
    Content     string
    Lines       string
    LineNumbers string
}
```

<a name="DataSet"></a>
## type [DataSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L4-L24>)



```go
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

    Distribution []FileResource
    // contains filtered or unexported fields
}
```

<a name="DataSource"></a>
## type [DataSource](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L73-L80>)



```go
type DataSource struct {
    FileObject *FileObject
    FileSet    *FileSet
    RecordSet  *RecordSet
    Extract    Extract
    Transform  Transform
    Format     Format
}
```

<a name="DataType"></a>
## type [DataType](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L50-L53>)



```go
type DataType struct {
    // MIME type
    DataType string
}
```

<a name="Extract"></a>
## type [Extract](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L82-L86>)



```go
type Extract struct {
    FileProperty ContentExtractionEnumeration
    Column       string
    JsonPath     string
}
```

<a name="Field"></a>
## type [Field](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L57-L64>)



```go
type Field struct {
    Source             Source
    Repeated           bool
    References         []*Field
    SubField           []*Field
    ParentField        []*Field
    EquivalentProperty string
}
```

<a name="FileObject"></a>
## type [FileObject](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L26-L28>)



```go
type FileObject struct {
    ContainedIn FileResource
}
```

<a name="FileResource"></a>
## type [FileResource](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L36-L39>)



```go
type FileResource struct {
    FileObject *FileObject
    FileSet    *FileSet
}
```

<a name="FileSet"></a>
## type [FileSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L30-L34>)



```go
type FileSet struct {
    ContainedIn FileResource
    Includes    string
    Excludes    string
}
```

<a name="Format"></a>
## type [Format](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L55>)



```go
type Format struct{}
```

<a name="JSON"></a>
## type [JSON](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L41>)



```go
type JSON string
```

<a name="RecordSet"></a>
## type [RecordSet](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L43-L48>)



```go
type RecordSet struct {
    Key      []*Field
    Data     []JSON
    Examples []JSON
    Source   Source
}
```

<a name="Source"></a>
## type [Source](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L66-L71>)



```go
type Source struct {
    DataSource *DataSource
    FileObject *FileObject
    FileSet    *FileSet
    RecordSet  *RecordSet
}
```

<a name="Split"></a>
## type [Split](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L94-L98>)



```go
type Split struct {
    TrainSplit      string
    TestSplit       string
    ValidationSplit string
}
```

<a name="Transform"></a>
## type [Transform](<https://github.com:b13rg/croissant-go/blob/main/pkg/croissant/types.go#L88-L92>)



```go
type Transform struct {
    Delimiter string
    Regex     string
    JsonQuery string
}
```