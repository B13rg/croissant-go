package croissant

import (
	"os"
	"path/filepath"
	"testing"
)

// Helper to create temporary Croissant test files
func createTestFile(t *testing.T, content string) string {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test_croissant.json")
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	return tmpFile
}

// Table-driven test cases for Croissant metadata validation
func TestValidateCroissantMetadataFiles(t *testing.T) {
	testCases := []struct {
		name      string
		json      string
		wantWarn  int
		wantError int
	}{
		{
			name: "Valid minimal dataset",
			json: `
{
	"@context": {"@language": "en"},
	"@type": "schema.org/Dataset",
	"name": "Test Dataset",
	"description": "A test dataset",
	"conformsTo": "http://mlcommons.org/croissant/1.0",
	"license": ["MIT"],
	"url": "http://example.com",
	"creator": ["Jane Doe"],
	"datePublished": "2024-01-01",
	"distribution": [],
	"recordSet": []
}
`,
			wantWarn:  2, // expects warnings for empty distribution and recordSet
			wantError: 0,
		},
		{
			name: "Missing required fields",
			json: `
{
	"@context": {"@language": "en"},
	"@type": "schema.org/Dataset",
	"conformsTo": "http://mlcommons.org/croissant/1.0",
	"description": "",
	"license": [],
	"name": "",
	"url": "",
	"creator": [],
	"datePublished": "",
	"distribution": [],
	"recordSet": []
}
`,
			wantWarn:  2, // distribution and recordSet warnings
			wantError: 7, // missing description, license, name, url, creator, datePublished
		},
		{
			name: "Invalid type and conformsTo",
			json: `
{
	"@context": {"@language": "en"},
	"@type": "SomethingElse",
	"name": "Dataset",
	"description": "desc",
	"conformsTo": "http://wrong.org/spec/1.2",
	"license": ["MIT"],
	"url": "http://example.com",
	"creator": ["Author"],
	"datePublished": "2024-01-01",
	"distribution": [],
	"recordSet": []
}
`,
			wantWarn:  2,
			wantError: 2, // wrong @type and conformsTo
		},
		{
			name: "Invalid version format",
			json: `
{
	"@context": {"@language": "en"},
	"@type": "schema.org/Dataset",
	"name": "Dataset",
	"description": "Test",
	"conformsTo": "http://mlcommons.org/croissant/1.0",
	"license": ["MIT"],
	"url": "http://example.com",
	"creator": ["John"],
	"datePublished": "2024-01-01",
	"version": "bad.version",
	"distribution": [],
	"recordSet": []
}
`,
			wantWarn:  3, // version warning + distribution + recordSet
			wantError: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			filePath := createTestFile(t, tc.json)
			ds, err := NewDataSetFromPath(filePath)
			if err != nil {
				t.Fatalf("Failed to load dataset: %v", err)
			}
			warns, errs := ds.Validate()

			if len(warns) != tc.wantWarn {
				t.Errorf("Expected %d warnings, got %d: %+v", tc.wantWarn, len(warns), warns)
			}
			if len(errs) != tc.wantError {
				t.Errorf("Expected %d errors, got %d: %+v", tc.wantError, len(errs), errs)
			}
		})
	}
}

// Optional: Test loading from a directory of croissant metadata files
func TestValidateAllCroissantFilesInDir(t *testing.T) {
	testDir := "testdata/croissant"
	files, err := os.ReadDir(testDir)
	if err != nil {
		t.Skip("Skipping directory test; testdata/croissant does not exist")
		return
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			t.Run(file.Name(), func(t *testing.T) {
				ds, err := NewDataSetFromPath(filepath.Join(testDir, file.Name()))
				if err != nil {
					t.Fatalf("Failed to load croissant file %s: %v", file.Name(), err)
				}
				_, errs := ds.Validate()
				if len(errs) > 0 {
					t.Errorf("Croissant file %s is invalid: %+v", file.Name(), errs)
				}
			})
		}
	}
}
