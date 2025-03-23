package util

import (
	"mime"
	"path/filepath"
	"strings"
)

func CompareList(phrase string, comparisionList []string) bool {
	for _, word := range comparisionList {
		if strings.Contains(phrase, word) {
			return true
		}
	}

	return false
}

// GetContentType returns the content type of a file based on its extension
func GetContentType(filename string) string {
	extension := filepath.Ext(filename)
	mimeType := mime.TypeByExtension(extension)

	return mimeType
}
