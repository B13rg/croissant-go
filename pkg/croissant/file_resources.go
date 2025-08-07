// Croissant spec filetypes and relations.
package croissant

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

func NewFileObject() *FileObject {
	return &FileObject{}
}

func (*FileObject) Validate() error {
	panic("not implemented")
}

func (*FileObject) ValidateHash() error {
	panic("not implemented")
}

// Update FileObject struct from resource
func (*FileObject) Update() error {
	panic("not implemented")
}

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
