# 🥐 croissant-go docs

A lightweight Go library for working with the Croissant data format, designed to provide formatting, validation, querying, and resource management capabilities.

- [Repo Readme](./README-repo.md)
- --- 
- [Official Spec](https://github.com/mlcommons/croissant/blob/main/docs/croissant-spec.md)

## Croissant Class Diagram

```mermaid
classDiagram
    Dataset -- "many" FileObject : distribution
    Dataset -- "many" FileSet : distribution
    Dataset -- "many" RecordSet : recordSet
    RecordSet -- "many" Field : field
    FileObject -- "many" FileSet : containedIn
    FileSet -- "many" FileObject : containedIn
    Field -- "one" DataSource : source
    DataSource -- "one" FileObject : fileObject
    DataSource -- "one" FileSet : fileSet
    DataSource -- "one" RecordSet : recordSet
    class Dataset {
        +@type
        +@context
        +name
        +description
        +license [ ]
        +url
        +creator [ ]
        +datePublished
        +keywords [ ]
        +publisher [ ]
        +version
        +dateCreated
        +dateModified
        +sameAs [ ]
        +sdLicense[ ]
        +inLanguage [ ]
        +distribution [ ]
        +dct:conformsTo
        +isLiveDataset
        +citeAs
    }
    class DataSource {
        +@id
        +fileObject  
        +fileSet
        +recordSet
        +extract
        +transform [ ]
        +format
    }
    class FileObject {
        +@id
        +sc:name
        +sc:contentUrl
        +sc:contentSize
        +sc:encodingFormat [ ]
        +sc:sameAs [ ]
        +sc:sha256
        +containedIn [ ]
    }
    class FileSet {
        +@id
        +containedIn [ ]
        +includes [ ]
        +excludes [ ]
        +sc:encodingFormat [ ]
    }
    class RecordSet {
        +@id
        +name
        +field [ ]
        +key [ ]
        +data [ ]
        +examples [ ]
        +annotation [ ]
    }
    class Field {
        +@id
        +name
        +source
        +dataType [ ]
        +isArray
        +arrayShape
        +references [ ]
        +subField [ ]
        +parentField [ ]
        +annotation [ ]
    }
    class Extract {
        +@id
        +fileProperty
        +column
        +jsonPath
    }
```
