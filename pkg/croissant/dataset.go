// Croissant spec filetypes and relations.
package croissant

import (
	"encoding/json"
	"os"

	"github.com/b13rg/croissant-go/pkg/types"
)

// [Dataset Class](https://docs.mlcommons.org/croissant/docs/croissant-spec.html#dataset-level-information)
// Based on https://docs.mlcommons.org/croissant/docs/croissant-spec.html#schemaorgdataset
type DataSet struct {
	// Must be `schema.org/Dataset`.
	NType string `json:"@type"`
	// Node ID
	NId string `json:"@id"`

	// Required Properties

	// Context alias definitions to make rest of document shorter.
	Context map[string]interface{} `json:"@context"`
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

	// Recommended Properties

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

	// Modified / Added Properties

	// List of FileObjects and FileSets associated with the dataset.
	// Modified from schema.org/Dataset.
	// Required.
	Distribution Distribution `json:"distribution,omitempty"`
	// Whether the dataset is a live dataset (in-process of being updated).
	IsLiveDataset bool `json:"isLiveDataset,omitempty"`
	// A citation to the dataset itself.
	CiteAs string `json:"citeAs,omitempty"`
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

func NewDataSetFromPath(filePath string) (*DataSet, error) {
	var dataSet DataSet
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return &dataSet, json.Unmarshal(data, &dataSet)
}

// Type used to group data resource objects together.
type DistributionItem interface{}

type Distribution struct {
	Items []DistributionItem
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
			d.Items = append(d.Items, fo)
		case "cr:FileSet":
			var fs FileSet
			if err := json.Unmarshal(raw, &fs); err != nil {
				return err
			}
			d.Items = append(d.Items, fs)
		default:
			return types.CroissantError{
				Message: "unknown @type",
				Value:   typeProbe.Type,
			}
		}
	}

	return nil
}

// The suggested context to use in a Croissant Json-LD file.
//
//nolint:gochecknoglobals
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
