// Croissant spec filetypes and relations.
package croissant

type FileObject struct {
	// Must be FileObject
	NType string `json:"@type"`
	// Node ID
	NId string `json:"@id"`
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

func NewFileObject() *FileObject {
	return &FileObject{}
}

func (*FileObject) Validate() error {
	panic("not implemented")
}

func (*FileObject) ValidateHash() error {
	panic("not implemented")
}

// Update FileObject struct from resource.
func (*FileObject) Update() error {
	panic("not implemented")
}

type FileSet struct {
	// Must be FileSet.
	NType string `json:"@type"`
	// Node ID
	NId string `json:"@id"`
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

func NewFileSet() *DataSet {
	return &DataSet{}
}

func (*FileSet) Validate() error {
	panic("not implemented")
}

// Update FileSet struct from resource.
func (*FileSet) Update() error {
	panic("not implemented")
}
