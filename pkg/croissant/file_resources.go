// Croissant spec filetypes and relations.
package croissant

type FileObject struct {
	NType          string        `json:"@type"`
	Name           string        `json:"sc:name"`
	ContentURL     string        `json:"sc:contentUrl"`
	ContentSize    string        `json:"sc:contentSize"`
	EncodingFormat string        `json:"sc:encodingFormat"`
	Sha256         string        `json:"sc:sha256,omitempty"`
	ContainedIn    *FileResource `json:"containedIn,omitempty"`
}

func NewFileObject() *FileObject {
	return &FileObject{}
}

func (*FileObject) Validate() error {
	panic("not implemented")
}

func (*FileObject) ValidateHash() error {
	panic("not implemented")
}

// Update Fileobject struct from resource
func (*FileObject) Update() error {
	panic("not implemented")
}

type FileSet struct {
	NType       string        `json:"@type"`
	ContainedIn *FileResource `json:"containedIn"`
	Includes    string        `json:"includes,omitempty"`
	Excludes    string        `json:"excludes,omitempty"`
}

func NewFileSet() *DataSet {
	return &DataSet{}
}

func (*FileSet) Validate() error {
	panic("not implemented")
}

// Update FileSet struct from resource
func (*FileSet) Update() error {
	panic("not implemented")
}

// Type used to group data resource objects together.
type FileResource interface {
	Validate() error
	Update() error
}
