package main

import (
	"path/filepath"
)

func findExtension(path string) string {
	ext := filepath.Ext(path)
	return ext
}
