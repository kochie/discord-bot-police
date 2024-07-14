package main

import (
	"path/filepath"
	"strings"
)

func compareList(phrase string, comparisionList []string) bool {
	for _, word := range comparisionList {
		if strings.Contains(phrase, word) {
			return true
		}
	}

	return false
}

func getContentType(filename string) string {
	extension := filepath.Ext(filename)
	switch extension {
	case ".jpg":
		return "image/jpeg"
	case ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".mp4":
		return "video/mp4"
	default:
		return ""
	}
}
