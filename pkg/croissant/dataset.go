// Croissant spec filetypes and relations.
package croissant

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

func NewDataSet() *DataSet {
	return &DataSet{
		Context:     SuggestedContext,
		NType:       "sc:DataSet",
		ConformsTo:  "http://mlcommons.org/croissant/1.0",
		Description: "",
		License:     []string{},
	}
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
