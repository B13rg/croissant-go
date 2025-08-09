package croissant

import (
	"os"
	"path/filepath"
	"testing"
)

// Helper to create temporary Croissant test files.
func createTestFile(t *testing.T, content string) string {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test_croissant.json")
	err := os.WriteFile(tmpFile, []byte(content), 0600)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	return tmpFile
}

// Table-driven test cases for Croissant metadata validation.
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
	"version": "1",
	"creator": ["Jane Doe"],
	"datePublished": "2024-01-01",
	"distribution": [],
	"recordSet": []
}
`,
			wantWarn:  2, // expects warnings for empty distribution and recordSet.
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
	"version": "",
	"creator": [],
	"datePublished": "",
	"distribution": [],
	"recordSet": []
}
`,
			wantWarn:  2, // distribution and recordSet.
			wantError: 6, // missing description, license, name, url, creator, datePublished.
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
			wantWarn:  2, // distribution and recordSet
			wantError: 2, // wrong @type and conformsTo.
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
			wantWarn:  3, // version warning + distribution + recordSet.
			wantError: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			filePath := createTestFile(t, testCase.json)
			ds, err := NewDataSetFromPath(filePath)
			if err != nil {
				t.Fatalf("Failed to load dataset: %v", err)
			}
			warns, errs := ds.Validate()

			if len(warns) != testCase.wantWarn {
				t.Errorf("Expected %d warnings, got %d: %+v", testCase.wantWarn, len(warns), warns)
			}
			if len(errs) != testCase.wantError {
				t.Errorf("Expected %d errors, got %d: %+v", testCase.wantError, len(errs), errs)
			}
		})
	}
}

// Optional: Test loading from a directory of croissant metadata files.
func TestValidateAllCroissantFilesInDir(t *testing.T) {
	testDir := "testdata/croissant/1.0"
	files, err := os.ReadDir(testDir)
	if err != nil {
		t.Skip("Skipping directory test; testdata/croissant does not exist")

		return
	}

	for _, dsDir := range files {
		if !dsDir.IsDir() {
			continue
		}
		dsFiles, err := os.ReadDir(filepath.Join(testDir, dsDir.Name()))
		if err != nil {
			t.Skip("Skipping directory test, error reading directory")

			return
		}
		for _, file := range dsFiles {
			if filepath.Ext(file.Name()) == ".json" {
				t.Run(file.Name(), func(t *testing.T) {
					ds, err := NewDataSetFromPath(filepath.Join(testDir, dsDir.Name(), file.Name()))
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
}
