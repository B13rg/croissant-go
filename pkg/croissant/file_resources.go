// Croissant spec filetypes and relations.
package croissant

import (
	"encoding/json"

	"github.com/b13rg/croissant-go/pkg/types"
)

// Type used to group data resource objects together.
type Distribution []DistributionItem
type DistributionItem interface {
	Validate() ([]types.CroissantWarning, []types.CroissantError)
}

func (d *Distribution) UnmarshalJSON(data []byte) error {
	// distribution can be an object or array of objects
	var rawItems []json.RawMessage

	if data[0] == '[' {
		// It's an array
		if err := json.Unmarshal(data, &rawItems); err != nil {
			return err
		}
	} else {
		// Single object, treat as one-element array
		rawItems = []json.RawMessage{data}
	}

	dist := []DistributionItem{}

	for _, raw := range rawItems {
		// Peek at @type
		var typeProbe struct {
			Type string `json:"@type"`
		}
		if err := json.Unmarshal(raw, &typeProbe); err != nil {
			return err
		}

		switch typeProbe.Type {
		case "cr:FileObject":
			var fo FileObject
			if err := json.Unmarshal(raw, &fo); err != nil {
				return err
			}
			dist = append(dist, &fo)
		case "cr:FileSet":
			var fs FileSet
			if err := json.Unmarshal(raw, &fs); err != nil {
				return err
			}
			dist = append(dist, &fs)
		default:
			return types.CroissantError{
				Message: "unknown @type",
				Value:   typeProbe.Type,
			}
		}
	}

	*d = dist

	return nil
}

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

func NewFileObject() *FileObject {
	return &FileObject{}
}

func (obj *FileObject) Validate() ([]types.CroissantWarning, []types.CroissantError) {
	listWarn := []types.CroissantWarning{}
	listError := []types.CroissantError{}

	if obj.Name == "" {
		listError = append(listError,
			types.CroissantError{
				Message: "resource name must be set",
				Value:   obj.ID,
			},
		)
	}

	if obj.ContentURL == "" {
		listError = append(listError,
			types.CroissantError{
				Message: "resource contentURL must be set",
				Value:   obj.ID,
			},
		)
	}

	if obj.ContentSize == "" {
		listWarn = append(listWarn,
			types.CroissantWarning{
				Message: "resource contentSize should be set",
				Value:   obj.ID,
			},
		)
	}

	if obj.EncodingFormat == "" {
		listWarn = append(listWarn,
			types.CroissantWarning{
				Message: "resource encodingFormat should be set",
				Value:   obj.ID,
			},
		)
	}

	if obj.Sha256 == "" {
		listWarn = append(listWarn,
			types.CroissantWarning{
				Message: "resource sha256 should be set",
				Value:   obj.ID,
			},
		)
	}

	return listWarn, listError
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

func NewFileSet() *DataSet {
	return &DataSet{}
}

func (obj *FileSet) Validate() ([]types.CroissantWarning, []types.CroissantError) {
	listWarn := []types.CroissantWarning{}
	if obj.Name == "" {
		listWarn = append(listWarn,
			types.CroissantWarning{
				Message: "resource name should be set",
				Value:   obj.ID,
			},
		)
	}
	if obj.Description == "" {
		listWarn = append(listWarn,
			types.CroissantWarning{
				Message: "resource description should be set",
				Value:   obj.ID,
			},
		)
	}
	if obj.EncodingFormat == "" {
		listWarn = append(listWarn,
			types.CroissantWarning{
				Message: "resource encodingFormat should be set",
				Value:   obj.ID,
			},
		)
	}

	return listWarn, []types.CroissantError{}
}

// Update FileSet struct from resource.
func (*FileSet) Update() error {
	panic("not implemented")
}
