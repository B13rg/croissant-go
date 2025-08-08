package util

import (
	"compress/gzip"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"os"
	"path/filepath"
	"sort"
)

// Walk a directory to build a list of all files in the tree.
func BuildDirFileList(directory string) ([]string, error) {
	clusterPaths := []string{}

	err := filepath.Walk(directory, func(path string, f os.FileInfo, err error) error {
		if f != nil && !f.IsDir() {
			clusterPaths = append(clusterPaths, path)
		}

		return nil
	})
	if err != nil {
		return []string{}, ErrorIfCheck("error building dir file list", err)
	}

	sort.Strings(clusterPaths)

	return clusterPaths, nil
}

// Write bytes to file (path included).
func WriteFile(input []byte, file string) error {
	f, err := os.Create(filepath.Clean(file))
	if err != nil {
		return err
	}
	defer f.Close()

	return os.WriteFile(file, input, 0600)
}

// Read bytes from file (path included).
func ReadFile(file string) ([]byte, error) {
	fCache, err := os.Open(filepath.Clean(file))
	if err != nil {
		return nil, err
	}
	defer fCache.Close()

	fileInfo, err := fCache.Stat()
	if err != nil {
		return nil, err
	}
	text := make([]byte, fileInfo.Size())
	_, err = fCache.Read(text)
	if err != nil {
		return nil, err
	}

	return text, nil
}

// Write bytes to gzip file (path included).
func WriteGzip(input []byte, file string) error {
	outFile, err := os.Create(filepath.Clean(file))
	if err != nil {
		return err
	}
	defer outFile.Close()
	gzipWriter := gzip.NewWriter(outFile)
	_, err = gzipWriter.Write(input)
	_ = gzipWriter.Close()

	return err
}

// Read bytes from a gzip file (path included).
func ReadGzip(filename string) ([]byte, error) {
	inFile, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return []byte{}, err
	}
	reader, err := gzip.NewReader(inFile)
	if err != nil {
		return []byte{}, err
	}
	data, err := io.ReadAll(reader)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

// Calculate the sha256 hash and returns the base64 encoded result.
func HashFile(path string) (string, error) {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return "", err
	}
	defer file.Close()

	hashBox := sha256.New()
	if _, err := io.Copy(hashBox, file); err != nil {
		return "", err
	}

	return base64.RawStdEncoding.EncodeToString(hashBox.Sum(nil)), nil
}
