// Croissant spec filetypes and relations.
package croissant

import (
	"encoding/json"
	"regexp"

	"github.com/b13rg/croissant-go/pkg/types"
	"github.com/b13rg/croissant-go/pkg/util"
)

// [Dataset Class](https://docs.mlcommons.org/croissant/docs/croissant-spec.html#dataset-level-information)
// Based on https://docs.mlcommons.org/croissant/docs/croissant-spec.html#schemaorgdataset
type DataSet struct {
	// Required Properties

	// Context alias definitions to make rest of document shorter.
	Context map[string]interface{} `json:"@context"`
	// Must be `schema.org/Dataset`.
	NType string `json:"@type"`
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

	// Whether the dataset is a live dataset (in-process of being updated).
	IsLiveDataset bool `json:"isLiveDataset,omitempty"`

	// List of FileObjects and FileSets associated with the dataset.
	// Modified from schema.org/Dataset.
	// Required.
	Distribution Distribution `json:"distribution,omitempty"`

	// List of RecordSets associated with the dataset
	RecordSets []RecordSet `json:"recordSet"`
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

func (ds *DataSet) Validate() ([]types.CroissantWarning, []types.CroissantError) {
	listWarn := []types.CroissantWarning{}
	listError := []types.CroissantError{}

	// Required Parameters
	if ds.NType != "schema.org/Dataset" {
		listError = append(listError,
			types.CroissantError{
				Message: "@type must be schema.org/Dataset",
				Value:   ds.NType,
			},
		)
	}

	if ds.ConformsTo != "http://mlcommons.org/croissant/1.0" &&
		ds.ConformsTo != "http://mlcommons.org/croissant/1.1" {
		listError = append(listError,
			types.CroissantError{
				Message: "conformsTo must be http://mlcommons.org/croissant/1.[0,1]",
				Value:   ds.ConformsTo,
			},
		)
	}

	if ds.Description == "" {
		listError = append(listError,
			types.CroissantError{
				Message: "dataset description should be set",
				Value:   "",
			},
		)
	}

	if len(ds.License) == 0 {
		listError = append(listError,
			types.CroissantError{
				Message: "dataset license should be set",
				Value:   "",
			},
		)
	}

	if ds.Name == "" {
		listError = append(listError,
			types.CroissantError{
				Message: "dataset name should be set",
				Value:   "",
			},
		)
	}

	if ds.URL == "" {
		listError = append(listError,
			types.CroissantError{
				Message: "dataset name should be set",
				Value:   "",
			},
		)
	}

	if len(ds.Creator) == 0 {
		listError = append(listError,
			types.CroissantError{
				Message: "dataset creator should be set",
				Value:   "",
			},
		)
	}

	if ds.DatePublished == "" {
		listError = append(listError,
			types.CroissantError{
				Message: "dataset datePublished should be set",
				Value:   "",
			},
		)
	}

	// Recommended Parameters

	// Start by checking for a plain num version.
	intVer := "^\\d+$"
	match, err := regexp.MatchString(intVer, ds.Version)
	cErr := util.ErrorIfCheck("error parsing version string", err)
	if cErr != nil {
		listError = append(listError, *cErr)
	}
	if !match {
		// https://semver.org/#is-there-a-suggested-regular-expression-regex-to-check-a-semver-string
		//nolint:lll
		semanticVer := "^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-((?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?$"
		match, err = regexp.MatchString(semanticVer, ds.Version)
		cErr := util.ErrorIfCheck("error parsing version string", err)
		if cErr != nil {
			listError = append(listError, *cErr)
		}
	}
	if !match {
		listWarn = append(listWarn,
			types.CroissantWarning{
				Message: "issue parsing dataset version as int or SemVer",
				Value:   ds.Version,
			},
		)
	}

	if len(ds.Distribution) == 0 {
		listWarn = append(listWarn,
			types.CroissantWarning{
				Message: "dataset distribution empty",
				Value:   "",
			},
		)
	}

	if len(ds.RecordSets) == 0 {
		listWarn = append(listWarn,
			types.CroissantWarning{
				Message: "dataset recordSets empty",
				Value:   "",
			},
		)
	}

	return listWarn, listError
}

func NewDataSetFromPath(filePath string) (*DataSet, error) {
	var dataSet DataSet
	data, err := util.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return &dataSet, json.Unmarshal(data, &dataSet)
}

func (ds *DataSet) WriteToFile(path string) error {
	str, err := json.Marshal(ds)
	if err != nil {
		return err
	}

	return util.WriteFile(str, path)
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
